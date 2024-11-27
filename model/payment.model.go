package model

import "github.com/google/uuid"

type Payment struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Paid   float64   `gorm:"type:decimal(10,2);not null" json:"paid" form:"paid" binding:"required"`
	Change float64   `gorm:"type:decimal(10,2);not null" json:"change" form:"change" binding:"required"`
	TransactionID uuid.UUID `gorm:"foreignKey:TransactionID" json:"transaction"`

	// Relationship
	Transaction	 Transaction
}
