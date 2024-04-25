package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"tenancy"
	"tenancy/mocks"
	"testing"
)

func TestMultiTenancyMiddleware(t *testing.T) {
	defTenant := &tenancy.TenantContext{Id: "default", Name: "Default"}
	tenant1 := &tenancy.TenantContext{Id: "tenant1", Name: "Tenant 01"}

	t.Run("Default tenant", func(t *testing.T) {
		tenantResolver := new(mocks.MockTenantResolver)
		tenantResolver.On("Resolve", mock.Anything, mock.Anything).Return(defTenant)

		rr := httptest.NewRecorder()

		r := gin.Default()
		r.ContextWithFallback = true
		r.Use(MultiTenancyMiddleware(tenantResolver))

		const url = "/test1"

		request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
		assert.NoError(t, err)

		r.GET(url, func(context *gin.Context) {
			currCtx := tenancy.FromContext(context)
			assert.Equal(t, currCtx, defTenant)
		})

		r.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("With tenant1", func(t *testing.T) {
		tenantResolver := new(mocks.MockTenantResolver)
		tenantResolver.On("Resolve", mock.Anything, defTenant.Id).Return(defTenant)
		tenantResolver.On("Resolve", mock.Anything, tenant1.Id).Return(tenant1)

		rr := httptest.NewRecorder()

		r := gin.Default()
		r.ContextWithFallback = true
		r.Use(MultiTenancyMiddleware(tenantResolver))

		const url = "/test2"

		request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
		request.Header.Add(tenancy.TenantHeader, tenant1.Id)
		assert.NoError(t, err)

		r.GET(url, func(context *gin.Context) {
			currCtx := tenancy.FromContext(context)
			assert.Equal(t, currCtx, tenant1)
		})

		r.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

}
