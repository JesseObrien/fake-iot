package storage

import "context"

type AccountPlanType string
type AccountPlanLimit int

const (
	AccountPlanTypeStandard   AccountPlanType = "standard"
	AccountPlanTypeEnterprise AccountPlanType = "enterprise"

	AccountPlanStandardLimit   AccountPlanLimit = 100
	AccountPlanEnterpriseLimit AccountPlanLimit = 1000
)

type AccountInfo struct {
	Id         string          `json:"id"`
	PlanType   AccountPlanType `json:"plan_type"`
	LoginCount int             `json:"login_count"`
	PlanLimit  int             `json:"plan_limit"`
}

type AccountStore interface {
	GetAccountInfo(ctx context.Context, accountId string) (AccountInfo, error)
	Upgrade(ctx context.Context, accountId string, plan AccountPlanType) error
}
