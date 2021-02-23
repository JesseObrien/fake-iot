package storage

import (
	"context"
	"time"
)

type UserLoginMetric struct {
	AccountID string    `json:"account_id"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

// AccountStore interface provides us a means of injecting our own stub for testing
type AccountStore interface {
	Write(ctx context.Context, metric UserLoginMetric) error
	CountByAccountId(ctx context.Context, accountId string) (int, error)
}
