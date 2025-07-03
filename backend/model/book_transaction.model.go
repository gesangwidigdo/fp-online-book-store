package model

import "github.com/google/uuid"

type BookTransaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"transaction_id" form:"transaction_id" binding:"required"`
	BookID        uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"book_id" form:"book_id" binding:"required"`
	Quantity      uint64    `gorm:"type:uint;default:0;not null" json:"quantity" form:"quantity" binding:"required"`
	Total         float64   `gorm:"type:decimal(10,2);default:0;not null" json:"total" form:"total" binding:"required"`

	Book        Book        `gorm:"foreignKey:BookID;references:ID" json:"book"`
	Transaction Transaction `gorm:"foreignKey:TransactionID;references:ID" json:"transaction"`
}
