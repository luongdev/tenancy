package gorm

import (
	"github.com/luongdev/tenancy/provider"
	"github.com/luongdev/tenancy/resolver"
	"gorm.io/gorm"
)

type DbProvider provider.DbProvider[*gorm.DB]

func NewDbProvider(cs ClientProvider, cr resolver.ConnectionStringResolver) DbProvider {
	return provider.NewDbProvider[*gorm.DB](cs, cr)
}
