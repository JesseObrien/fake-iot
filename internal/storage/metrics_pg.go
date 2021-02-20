package storage

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type PgAccountStore struct {
	conn *sql.DB
}

func NewPgAccountStore(conn *sql.DB) *PgAccountStore {
	return &PgAccountStore{conn}
}

func (pas *PgAccountStore) Write(ctx context.Context, metric UserLoginMetric) error {
	sql := `INSERT INTO account_logins (user_id, account_id, timestamp) VALUES ($1, $2, $3)`

	_, err := pas.conn.QueryContext(ctx, sql, metric.UserID, metric.AccountID, metric.Timestamp)

	if err != nil {
		return fmt.Errorf("error inserting metric into database: %v", err)
	}

	return nil
}

func (pas *PgAccountStore) CountByAccountId(ctx context.Context, accountId string) (int, error) {
	sql := `SELECT COUNT(*) FROM account_logins WHERE account_id=$1`

	var count int
	err := pas.conn.QueryRowContext(ctx, sql, accountId).Scan(&count)

	if err != nil {
		return count, fmt.Errorf("could not execute sql statement %w", err)
	}

	return count, nil
}
