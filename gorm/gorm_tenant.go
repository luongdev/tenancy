package gorm

type TenantEntity struct {
	TenantId string `gorm:"index;column:tenantId"`
}
