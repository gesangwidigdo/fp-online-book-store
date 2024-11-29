package model

import "github.com/google/uuid"

type Transaction struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	GrandTotal      float64   `gorm:"type:decimal(10,2);default:0" json:"grand_total" form:"grand_total" binding:"required"`
	TransactionTime string    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"transaction_time" form:"transaction_time" binding:"required"`
	
	// Relationship
	Books           []Book    `gorm:"many2many:book_transactions"`
	UserID          uint64    `gorm:"foreignKey:UserID" json:"user"`
}
