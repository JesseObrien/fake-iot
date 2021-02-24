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
	mu               *sync.Mutex
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
}

func NewAccountUpdateStore() *AccountUpdateStore {
	return &AccountUpdateStore{
		map[string][]AccountUpdateSubscription{},
	}
}

func (aus *AccountUpdateStore) Subscribe(accountId string) *AccountUpdateSubscription {
	updates := aus.accountUpdates[accountId]
	newSubscription := AccountUpdateSubscription{
		SubscriptionReference: uuid.New().String(),
		Updates:               make(chan AccountUpdate),
		subscriptionList:      &updates,
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
