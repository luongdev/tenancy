package resolver

import (
	"context"
	"tenancy"
)

type TenantResolver interface {
	Resolve(ctx context.Context, id string) *tenancy.TenantContext
}

type ConnectionStringResolver interface {
	Resolve(ctx context.Context, id string) (string, error)
}
