package repository

import "gorm.io/gorm"

type PaymentRepository interface {

}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return paymentRepository{
		DB : db,
	}
}

