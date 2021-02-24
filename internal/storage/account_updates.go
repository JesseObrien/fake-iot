package storage

import (
	"sync"

	"github.com/google/uuid"
)

type AccountUpdate struct {
	AccountId  string `json:"account_id"`
	LoginCount int    `json:"login_count"`
}

type AccountUpdateSubscription struct {
	SubscriptionReference string
	Updates               chan AccountUpdate

	// the list of subscriptions this subscription belongs to. used for unsubscribe
	subscriptionList *[]AccountUpdateSubscription

	// Link to the top level mutex in the update store to lock when we're unsubscribing
	mu *sync.Mutex
}

func (aus *AccountUpdateSubscription) Unsubscribe() {
	aus.mu.Lock()
	defer aus.mu.Unlock()

	for idx, subscription := range *aus.subscriptionList {
		if subscription.SubscriptionReference == aus.SubscriptionReference {
			// Close the updates channel
			close(subscription.Updates)
			// Remove the subscription from the list and break out
			s := *aus.subscriptionList
			s = append(s[:idx], s[idx+1:]...)
			*aus.subscriptionList = s
			break
		}
	}
}

type AccountUpdateStore struct {
	accountUpdates map[string][]AccountUpdateSubscription
	mu             *sync.Mutex
}

func NewAccountUpdateStore() *AccountUpdateStore {
	return &AccountUpdateStore{
		map[string][]AccountUpdateSubscription{},
		&sync.Mutex{},
	}
}

func (aus *AccountUpdateStore) Subscribe(accountId string) *AccountUpdateSubscription {
	aus.mu.Lock()
	defer aus.mu.Unlock()
	updates := aus.accountUpdates[accountId]
	newSubscription := AccountUpdateSubscription{
		SubscriptionReference: uuid.New().String(),
		Updates:               make(chan AccountUpdate),
		subscriptionList:      &updates,
		mu:                    aus.mu,
	}
	aus.accountUpdates[accountId] = append(updates, newSubscription)
	return &newSubscription
}

func (aus *AccountUpdateStore) Fanout(accountId string, update AccountUpdate) {
	// Get the subscriptions, if there are none, the range will no-op
	subscriptions := aus.accountUpdates[accountId]

	for _, subscription := range subscriptions {
		subscription.Updates <- update
	}
}
