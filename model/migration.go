package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&Book{},
		&Transaction{},
		&BookTransaction{},
	)
	if err != nil {
		return err
	}

	return nil
}
