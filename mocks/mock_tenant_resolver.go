package mocks

import (
	"context"
	"github.com/luongdev/tenancy"
	"github.com/stretchr/testify/mock"
)

type MockTenantResolver struct {
	mock.Mock
}

func (m *MockTenantResolver) Resolve(ctx context.Context, id string) *tenancy.TenantContext {
	ret := m.Called(ctx, id)

	var r0 *tenancy.TenantContext
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*tenancy.TenantContext)
	}

	return r0
}
