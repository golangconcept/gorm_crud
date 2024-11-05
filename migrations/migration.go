// migrations/migration.go
package migrations

import (
	"fmt"
	"gorm_crud/models"

	"gorm.io/gorm"
)

// RunMigrations handles the database schema migrations
func RunMigrations(db *gorm.DB) error {
	// Automatically migrate the schema for the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}
	return nil
}
