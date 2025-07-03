package seeder

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Djuanzz/go-template/model"
	"gorm.io/gorm"
)

func BookSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migration/json/book.seed.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var books []model.Book
	if err := json.Unmarshal(jsonData, &books); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&model.Book{})
	if !hasTable {
		if err := db.AutoMigrate(&model.Book{}); err != nil {
			return err
		}
	}

	for _, data := range books {
		var book model.Book
		err := db.Where(&model.Book{ISBN: data.ISBN}).First(&book).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&book, "isbn = ? OR slug = ?", data.ISBN, data.Slug).RowsAffected

		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}


	return nil
}