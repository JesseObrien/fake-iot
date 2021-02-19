package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgAccountStore struct {
	conn *pgxpool.Pool
}

func NewPgAccountStore(conn *pgxpool.Pool) *PgAccountStore {
	return &PgAccountStore{conn}
}

func (pas *PgAccountStore) Write(ctx context.Context, metric UserLoginMetric) error {
	tx, err := pas.conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return fmt.Errorf("could not begin transaction %w", err)
	}

	sql := `INSERT INTO account_logins (user_id, account_id, timestamp) VALUES ($1, $2, $3)`

	_, err = tx.Exec(ctx, sql, metric.UserID, metric.AccountID, metric.Timestamp)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			log.Printf("error rolling back transaction %v", err)
		}

		log.Printf("database failure: %v", err)
		return fmt.Errorf("could not execute sql statement %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("could not commit sql transaction %w", err)
	}

	return nil
}

func (pas *PgAccountStore) CountByAccountId(ctx context.Context, accountId string) (int, error) {

	sql := `SELECT COUNT(*) FROM account_logins WHERE account_id=$1`

	var count int
	err := pas.conn.QueryRow(ctx, sql, accountId).Scan(&count)

	if err != nil {
		return count, fmt.Errorf("could not execute sql statement %w", err)
	}

	return count, nil
}
