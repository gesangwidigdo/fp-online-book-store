package seeder

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MustHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func UserSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migration/json/user.seed.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var users []model.User
	if err := json.Unmarshal(jsonData, &users); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&model.User{})
	if !hasTable {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			return err
		}
	}

	for _, data := range users {
		var user model.User
		err := db.Where(&model.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ? OR name = ?", data.Email, data.Name).RowsAffected

		password, err := utils.HashPassword(data.Password)
		if err != nil {
			return err
		}
		
		data.Password = password
		
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}