package provider

import "context"

type ClientProvider[TClient interface{}] interface {
	Get(ctx context.Context, dsn string) (TClient, error)
}

type DbProvider[TClient interface{}] interface {
	Get(ctx context.Context, tenantId string) (TClient, error)
}
