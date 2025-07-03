package migration

import (
	"github.com/Djuanzz/go-template/migration/seeder"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeder.UserSeeder(db); err != nil {
		return err
	}

	if err := seeder.BookSeeder(db); err != nil {
		return err
	}

	return nil
}