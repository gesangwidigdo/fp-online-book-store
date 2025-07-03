package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	GrandTotal float64   `gorm:"type:decimal(10,2);default:0" json:"grand_total" form:"grand_total" binding:"required"`
	Status     string    `gorm:"type:varchar(50);default:draft" json:"status"`
	UserID     uuid.UUID `gorm:"foreignKey:UserID" json:"user"`

	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`

	// Relationship
	BookTransaction []BookTransaction `gorm:"foreignKey:TransactionID" json:"book_transaction"`
}
