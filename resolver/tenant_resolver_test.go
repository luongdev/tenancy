package resolver

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tenancy"
	"tenancy/mocks"
	"testing"
)

func TestTenantResolver(t *testing.T) {
	tenantResolver := new(mocks.MockTenantResolver)

	defTenant := &tenancy.TenantContext{Id: "default", Name: "Default"}

	tenantResolver.On("Resolve", mock.Anything, defTenant.Id).Return(defTenant)

	tenant := tenantResolver.Resolve(context.Background(), defTenant.Id)

	assert.Equal(t, defTenant, tenant)
}
