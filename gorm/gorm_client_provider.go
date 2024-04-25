package gorm

import (
	"context"
	"gorm.io/gorm"
	"tenancy/provider"
)

type ClientProvider provider.ClientProvider[*gorm.DB]

type ClientProviderFunc provider.ClientProviderFunc[*gorm.DB]

func (c ClientProviderFunc) Get(ctx context.Context, dsn string) (*gorm.DB, error) {
	return c(ctx, dsn)
}
