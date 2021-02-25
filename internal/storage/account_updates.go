package storage

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type AccountUpdate struct {
	AccountId  string `json:"account_id"`
	LoginCount int    `json:"login_count"`
}

type AccountUpdateSubscription struct {
	AccountId             string
	SubscriptionReference string
	Updates               chan AccountUpdate
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
		AccountId:             accountId,
		SubscriptionReference: uuid.New().String(),
		Updates:               make(chan AccountUpdate, 100),
	}
	aus.accountUpdates[accountId] = append(updates, newSubscription)
	return &newSubscription
}

func (aus *AccountUpdateStore) Unsubscribe(subscription *AccountUpdateSubscription) error {
	aus.mu.Lock()
	defer aus.mu.Unlock()

	subscriptionList, ok := aus.accountUpdates[subscription.AccountId]
	if !ok {
		return fmt.Errorf("account subscriptions do not exist for %v", subscription.AccountId)
	}

	for idx, sub := range subscriptionList {
		if subscription.SubscriptionReference == sub.SubscriptionReference {
			// Close the updates channel
			close(subscription.Updates)
			// Remove the subscription from the list and break out
			s := subscriptionList
			s = append(s[:idx], s[idx+1:]...)
			aus.accountUpdates[subscription.AccountId] = s
			break
		}
	}

	return nil
}

func (aus *AccountUpdateStore) Fanout(accountId string, update AccountUpdate) {
	aus.mu.Lock()
	defer aus.mu.Unlock()

	// Get the subscriptions, if there are none, the range will no-op
	subscriptions := aus.accountUpdates[accountId]

	for _, subscription := range subscriptions {
		select {
		case subscription.Updates <- update: // Write to the channel unless it's blocked
		default: // if the channel is blocked, no-op
		}
	}
}
