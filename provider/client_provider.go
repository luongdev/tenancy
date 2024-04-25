package provider

import "context"

type ClientProviderFunc[TClient interface{}] func(ctx context.Context, dsn string) (TClient, error)

func (c ClientProviderFunc[TClient]) Get(ctx context.Context, dsn string) (TClient, error) {
	return c(ctx, dsn)
}
