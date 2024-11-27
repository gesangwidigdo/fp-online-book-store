package model

import "github.com/google/uuid"

type Book struct {
	ID              uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	ISBN            string        `gorm:"type:varchar(20);unique;not null" json:"isbn" form:"isbn" binding:"required"`
	Title           string        `gorm:"type:varchar(255);not null" json:"title" form:"title" binding:"required"`
	Slug            string        `gorm:"type:varchar(255);unique;not null" json:"slug" form:"slug" binding:"required"`
	Author          string        `gorm:"type:varchar(255);not null" json:"author" form:"author" binding:"required"`
	Summary         string        `gorm:"type:text;not null" json:"summary" form:"summary" binding:"required"`
	BookImage       string        `gorm:"type:text" json:"book_image" form:"book_image" binding:"required"`
	PublicationYear string        `gorm:"type:varchar(4)" json:"publication_year" form:"publication_year" binding:"required"`
	Price           float64       `gorm:"type:decimal(10,2);default:0;not null" json:"price" form:"price" binding:"required"`

	// Relationship
	Transactions    []Transaction `gorm:"many2many:book_transactions;"`
}
