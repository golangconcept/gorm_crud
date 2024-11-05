package migrations

import (
	"gorm_crud/models" // import your models package

	"gorm.io/gorm"
)

// Migration to create the products table
func MigrateCreateProducts(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})
}
