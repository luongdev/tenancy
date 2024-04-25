package tenancy

import "context"

const TenantHeader = "X-Tenant"
const CurrentTenantKey = "current_tenant"

type TenantContext struct {
	Id   string
	Name string
}

func _defaultTenantContext() *TenantContext {
	return &TenantContext{Id: "default", Name: "Default"}
}

func CurrentTenant(ctx context.Context, id, name string) context.Context {
	if id == "" {
		return context.WithValue(ctx, CurrentTenantKey, _defaultTenantContext())
	}
	return context.WithValue(ctx, CurrentTenantKey, &TenantContext{Id: id, Name: name})
}

func FromContext(ctx context.Context) *TenantContext {
	val := ctx.Value(CurrentTenantKey)
	if val == nil {
		return _defaultTenantContext()
	}

	return val.(*TenantContext)
}
