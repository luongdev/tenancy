package resolver

import (
	"context"
	"github.com/luongdev/tenancy"
)

type TenantResolver interface {
	Resolve(ctx context.Context, id string) *tenancy.TenantContext
}

type ConnectionStringResolver interface {
	Resolve(ctx context.Context, id string) (string, error)
}
