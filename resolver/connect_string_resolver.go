package resolver

import (
	"context"
	"tenancy"
)

type ConnectionStringResolveFunc func(ctx context.Context, id string) (string, error)

func (c ConnectionStringResolveFunc) Resolve(ctx context.Context, key string) (string, error) {
	return c(ctx, key)
}

var _ ConnectionStringResolver = (*ConnectionStringResolveFunc)(nil)

func ConnectionStringResolvers(resolvers ...ConnectionStringResolver) ConnectionStringResolver {
	return ConnectionStringResolveFunc(func(ctx context.Context, id string) (string, error) {
		for _, resolver := range resolvers {
			result, err := resolver.Resolve(ctx, id)
			if err != nil {
				return "", err
			}
			if len(result) > 0 {
				return result, nil
			}
		}
		return "", nil
	})
}

type DefaultConnectionStringResolver struct {
	fallback ConnectionStringResolver
}

func NewConnectionStringResolver(fallback ConnectionStringResolver) *DefaultConnectionStringResolver {
	return &DefaultConnectionStringResolver{
		fallback: fallback,
	}
}

func (m *DefaultConnectionStringResolver) Resolve(ctx context.Context, key string) (string, error) {
	tenant := tenancy.FromContext(ctx)
	if tenant.Id == "default" {
		return m.fallback.Resolve(ctx, key)
	}

	return m.fallback.Resolve(ctx, tenant.Id)
}
