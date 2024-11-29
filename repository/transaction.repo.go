package repository

import (
	"github.com/Djuanzz/go-template/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(id string) (model.Transaction, error)
	FindTransactionStatusByUserID(id string) (model.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: db,
	}
}

func (tr *transactionRepository) Create(id string) (model.Transaction, error) {
	var transaction model.Transaction
	userId, _ := uuid.Parse(id)
	transaction.UserID = userId

	if err := tr.DB.Create(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (tr *transactionRepository) FindTransactionStatusByUserID(id string) (model.Transaction, error) {
	var transaction model.Transaction
	userId, _ := uuid.Parse(id)

	if err := tr.DB.Where("user_id = ? AND status = ?", userId, false).First(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}