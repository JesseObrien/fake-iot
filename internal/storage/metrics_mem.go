package storage

import (
	"context"
	"fmt"
	"sync"
)

// MemMetricStore is a memory implementation of the account store
// to use for testing purposes, etc.
type MemMetricStore struct {
	metrics map[string][]UserLoginMetric
	mu      sync.Mutex
}

func NewMemMetricStore() *MemMetricStore {
	return &MemMetricStore{
		metrics: map[string][]UserLoginMetric{},
	}
}

func (mas *MemMetricStore) WroteMetric(accountId, userId string) (bool, error) {
	mas.mu.Lock()
	defer mas.mu.Unlock()
	metrics, ok := mas.metrics[accountId]

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

func (mas *MemMetricStore) Write(ctx context.Context, metric UserLoginMetric) error {
	mas.mu.Lock()
	defer mas.mu.Unlock()

	metrics := mas.metrics[metric.AccountID]
	mas.metrics[metric.AccountID] = append(metrics, metric)

	return nil
}

func (mas *MemMetricStore) CountByAccountId(ctx context.Context, accountId string) (int, error) {
	mas.mu.Lock()
	defer mas.mu.Unlock()

	return len(mas.metrics[accountId]), nil
}
