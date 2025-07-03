package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name" form:"name" binding:"required"`
	Email       string    `gorm:"type:varchar(255);unique;not null" json:"email" form:"email" binding:"required"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password" form:"password" binding:"required"`
	Address     string    `gorm:"type:text;not null" json:"address" form:"address" binding:"required"`
	Gender      string    `gorm:"type:varchar(25);not null" json:"gender" form:"gender" binding:"required"`
	PhoneNumber string    `gorm:"type:varchar(15);unique;not null" json:"phone_number" form:"phone_number" binding:"required"`
	BirthDate   string    `gorm:"type:date;not null" json:"birth_date" form:"birth_date" binding:"required"`
	UserType    string    `gorm:"type:varchar(25);default:user;not null" json:"user_type" form:"user_type" binding:"required"`

	// Relationship
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
}
