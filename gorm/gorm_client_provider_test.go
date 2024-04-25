package gorm

import (
	"context"
	"database/sql"
	"github.com/luongdev/tenancy/provider"
	"github.com/luongdev/tenancy/resolver"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestClientProvider(t *testing.T) {
	clientProvider := ClientProviderFunc(func(ctx context.Context, dsn string) (*gorm.DB, error) {
		var client *gorm.DB
		db, err := sql.Open("sqlite3", dsn)
		if err != nil {
			return nil, err
		}

		db.SetMaxIdleConns(1)
		db.SetMaxOpenConns(1)

		client, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

		return client.WithContext(ctx), err
	})

	connResolver := resolver.NewConnectionStringResolver(resolver.ConnectionStringResolveFunc(func(ctx context.Context, id string) (string, error) {
		return "./test.db", nil
	}))

	dbProvider := provider.NewDbProvider[*gorm.DB](clientProvider, connResolver)
	db, _ := dbProvider.Get(context.Background(), "default")

	type Tenant struct {
		TenantEntity

		Id   string `gorm:"type:varchar(36);primary_key"`
		Name string `gorm:"type:varchar(255);not null"`
	}

	tx := db.Exec("create table tenants (id varchar(36) primary key, name varchar(255) not null, tenantId varchar(36))")
	log.Println(tx)

	res := db.Create(&Tenant{Name: "Test"})

	tenant := &Tenant{}
	res = db.Find(tenant, "name = ?", "Test")

	log.Println(tenant)

	if res.Error != nil {
		t.Fatalf("Error creating tenant: %v", res.Error)
	}

	log.Println("Tenant created successfully")

	//db, err := clientProvider.Get(tenancy.CurrentTenant(context.Background(), "default", "Default"), "./test.db")
	//if err != nil {
	//	t.Fatalf("Error getting db: %v", err)
	//}
	//
	//type Tenant struct {
	//	TenantEntity
	//
	//	Id   string `gorm:"type:varchar(36);primary_key"`
	//	Name string `gorm:"type:varchar(255);not null"`
	//}
	//
	//tx := db.Exec("create table tenants (id varchar(36) primary key, name varchar(255) not null, tenant_id varchar(36))")
	//log.Println(tx)
	//
	//res := db.Create(&Tenant{Name: "Test"})
	//
	//tenant := &Tenant{}
	//res = db.Find(tenant, "name = ?", "Test")
	//
	//log.Println(tenant)
	//
	//if res.Error != nil {
	//	t.Fatalf("Error creating tenant: %v", res.Error)
	//}
	//
	//log.Println("Tenant created successfully")
}
