package provider

import (
	"context"
	"github.com/luongdev/tenancy/resolver"
)

type DefaultDbProvider[TClient interface{}] struct {
	clientProvider           ClientProvider[TClient]
	connectionStringResolver resolver.ConnectionStringResolver
}

func NewDbProvider[TClient interface{}](provider ClientProvider[TClient], resolver resolver.ConnectionStringResolver) (d *DefaultDbProvider[TClient]) {
	d = &DefaultDbProvider[TClient]{
		clientProvider:           provider,
		connectionStringResolver: resolver,
	}
	return
}

func (d *DefaultDbProvider[TClient]) Get(ctx context.Context, tenantId string) (TClient, error) {
	connStr, err := d.connectionStringResolver.Resolve(ctx, tenantId)
	if err != nil {
		panic(err)
	}

	client, err := d.clientProvider.Get(ctx, connStr)
	if err != nil {
		panic(err)
	}

	return client, nil
}
