package gin

import (
	"github.com/gin-gonic/gin"
	"tenancy"
	"tenancy/resolver"
)

func MultiTenancyMiddleware(resolver resolver.TenantResolver) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenantHeader := ctx.Request.Header.Get(tenancy.TenantHeader)
		tenant := resolver.Resolve(ctx, tenantHeader)

		tenantCtx := tenancy.CurrentTenant(ctx, tenant.Id, tenant.Name)
		ctx.Request = ctx.Request.WithContext(tenantCtx)

		ctx.Next()
	}
}
