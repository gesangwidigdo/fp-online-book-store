package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&Book{},
		&Transaction{},
		&Payment{},
		&BookTransaction{},
	)
	if err != nil {
		return err
	}

	return nil
}
