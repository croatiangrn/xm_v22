package database

import (
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&company.Company{})
	if err != nil {
		return err
	}
	return nil
}
