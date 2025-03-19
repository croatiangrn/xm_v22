package database

import (
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"gorm.io/gorm"
)

func createEnumTypes(db *gorm.DB) error {
	// Check if the enum type already exists
	var exists bool
	err := db.Raw(`
        SELECT EXISTS (
            SELECT 1 FROM pg_type WHERE typname = 'company_type'
        );
    `).Scan(&exists).Error
	if err != nil {
		return err
	}

	// If the enum type doesn't exist, create it
	if !exists {
		err = db.Exec(`
            CREATE TYPE company_type AS ENUM (
                'corporations', 'non_profit', 'cooperative', 'sole_proprietorship'
            );
        `).Error
		if err != nil {
			return err
		}
	}

	return nil
}
func AutoMigrate(db *gorm.DB) error {
	if err := createEnumTypes(db); err != nil {
		return err
	}

	if err := db.AutoMigrate(&company.Company{}); err != nil {
		return err
	}

	return nil
}
