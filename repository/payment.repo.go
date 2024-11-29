package repository

import (
	"github.com/Djuanzz/go-template/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (pr paymentRepository) Create(payment model.Payment) (model.Payment, error) {
		if err := pr.DB.Create(&payment).Error; err != nil {
			return model.Payment{}, err
		}
	
	return payment, nil
}

func (pr paymentRepository) UpdateStatusByTranscationID(transId uuid.UUID) (model.Payment, error) {
	var payment model.Payment
	// if err := pr.DB.Where("transaction_id = ?", transId).Take(&payment).Error; err != nil {
	// 	return model.Payment{}, err
	// }

	// if err := pr.DB.Where("transaction_id = ?", transId).Update("status", )

	return payment, nil
}