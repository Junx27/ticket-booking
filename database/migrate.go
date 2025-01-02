package database

import (
	"fmt"

	"github.com/Junx27/ticket-booking/entity"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {

	if err := db.Migrator().DropTable(&entity.User{}); err != nil {
		return fmt.Errorf("failed to drop table: %w", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
