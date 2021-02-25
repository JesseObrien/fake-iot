package storage

import (
	"context"
	"database/sql"
	"fmt"
)

type PgAccountStore struct {
	conn        *sql.DB
	metricStore MetricStore
}

func NewPgAccountStore(conn *sql.DB, metricStore MetricStore) AccountStore {
	return &PgAccountStore{conn, metricStore}
}

func (pas *PgAccountStore) GetAccountInfo(ctx context.Context, accountId string) (AccountInfo, error) {
	info := AccountInfo{}

	sql := `SELECT id, plan_type FROM accounts WHERE id=$1`

	err := pas.conn.QueryRowContext(ctx, sql, accountId).Scan(&info.Id, &info.PlanType)

	if err != nil {
		return info, fmt.Errorf("error retrieving account info: %w", err)
	}

	switch info.PlanType {
	case AccountPlanTypeStandard:
		info.PlanLimit = AccountPlanStandardLimit
	case AccountPlanTypeEnterprise:
		info.PlanLimit = AccountPlanEnterpriseLimit
	default:
		return info, fmt.Errorf("plan type %s is not valid", info.PlanType)
	}

	loginCount, err := pas.metricStore.CountByAccountId(ctx, accountId)
	if err != nil {
		return info, fmt.Errorf("error retrieving login counts: %w", err)
	}

	info.LoginCount = loginCount

	return info, nil
}

func (pas *PgAccountStore) Upgrade(ctx context.Context, accountId string, planType AccountPlanType) error {
	sql := `UPDATE accounts SET plan_type=$1 WHERE id=$2`

	_, err := pas.conn.QueryContext(ctx, sql, planType, accountId)

	if err != nil {
		return fmt.Errorf("error upgrading account: %w", err)
	}

	return nil
}
