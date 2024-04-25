package provider

import (
	"context"
	"tenancy/resolver"
)

type DefaultDbProvider[TClient interface{}] struct {
	clientProvider           ClientProvider[TClient]
	connectionStringResolver resolver.ConnectionStringResolver
}

func NewDbProvider[TClient interface{}](provider ClientProvider[TClient], resolver resolver.ConnectionStringResolver) *DefaultDbProvider[TClient] {
	return &DefaultDbProvider[TClient]{
		clientProvider:           provider,
		connectionStringResolver: resolver,
	}
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
