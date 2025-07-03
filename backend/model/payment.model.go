package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" form:"id" binding:"required"`
	Method string `gorm:"type:varchar(50);null;default:bca" json:"method" form:"method" binding:"required"`
	Amount int64 `gorm:"type:bigint;not null" json:"amount" form:"amount" binding:"required"`
	Status string `gorm:"type:varchar(50);default:pending" json:"status" form:"status" binding:"required"`
	Date	time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"date" form:"date"`
	TransactionID uuid.UUID    `gorm:"type:uuid;null" json:"transaction_id"`
	
	// Relationship
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID" json:"transaction"`
}
