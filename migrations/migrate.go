package migrations

import (
	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	sqlPath := "migrations/database.up.sql"

	sqlContent, err := os.ReadFile(sqlPath)
	if err != nil {
		return err
	}

	if err := db.Exec(string(sqlContent)).Error; err != nil {
		return err
	}

	return nil
}

func Down(db *gorm.DB) error {
	sqlPath := "migrations/database.down.sql"

	sqlContent, err := os.ReadFile(sqlPath)
	if err != nil {
		return err
	}

	if err := db.Exec(string(sqlContent)).Error; err != nil {
		return err
	}

	return nil
}
