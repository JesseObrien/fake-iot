package storage

import (
	"context"
	"fmt"
	"sync"
)

// MemAccountStore is a memory implementation of the account store
// to use for testing purposes, etc.
type MemAccountStore struct {
	metrics map[string][]UserLoginMetric
	mu      sync.Mutex
}

func NewMemAccountStore() *MemAccountStore {
	return &MemAccountStore{
		metrics: map[string][]UserLoginMetric{},
	}
}

func (mas *MemAccountStore) WroteMetric(accountId, userId string) (bool, error) {
	mas.mu.Lock()
	metrics, ok := mas.metrics[accountId]
	mas.mu.Unlock()

	if !ok {
		return false, fmt.Errorf("no metrics found for account")
	}

	for _, metric := range metrics {
		if metric.UserID == userId {
			return true, nil
		}
	}

	return false, fmt.Errorf("could not find write for account id: %s, user_id: %s", accountId, userId)
}

func (mas *MemAccountStore) Write(ctx context.Context, metric UserLoginMetric) error {
	mas.mu.Lock()
	defer mas.mu.Unlock()
	if _, ok := mas.metrics[metric.AccountID]; !ok {
		mas.metrics[metric.AccountID] = []UserLoginMetric{metric}
		return nil
	}

	metrics := mas.metrics[metric.AccountID]
	mas.metrics[metric.AccountID] = append(metrics, metric)
}

func (mas *MemAccountStore) CountByAccountId(ctx context.Context, accountId string) (int, error) {
	mas.mu.Lock()
	defer mas.mu.Unlock()

	if _, ok := mas.metrics[accountId]; !ok {
		return 0, nil
	}

	return len(mas.metrics[accountId]), nil
}
