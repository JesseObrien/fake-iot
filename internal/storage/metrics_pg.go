package storage

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type PgMetricStore struct {
	conn *sql.DB
}

func NewPgMetricStore(conn *sql.DB) MetricStore {
	return &PgMetricStore{conn}
}

func (pms *PgMetricStore) Write(ctx context.Context, metric UserLoginMetric) error {
	sql := `INSERT INTO account_logins (user_id, account_id, timestamp) VALUES ($1, $2, $3)`

	_, err := pms.conn.QueryContext(ctx, sql, metric.UserID, metric.AccountID, metric.Timestamp)

	if err != nil {
		return fmt.Errorf("error inserting metric into database: %w", err)
	}

	return nil
}

func (pms *PgMetricStore) CountByAccountId(ctx context.Context, accountId string) (int, error) {
	sql := `SELECT COUNT(*) FROM account_logins WHERE account_id=$1`

	var count int
	err := pms.conn.QueryRowContext(ctx, sql, accountId).Scan(&count)

	if err != nil {
		return count, fmt.Errorf("could not execute sql statement %w", err)
	}

	return count, nil
}
