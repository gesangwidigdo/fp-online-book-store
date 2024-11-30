package repository

import (
	"fmt"

	"github.com/Djuanzz/go-template/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(id string) (model.Transaction, error)
	FindTransactionStatusByUserID(id string, status string) (model.Transaction, error)
	GetAllTransactionByUserLogin(id string) ([]model.Transaction, error)
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

func (tr *transactionRepository) FindTransactionStatusByUserID(id string, status string) (model.Transaction, error) {
	var transaction model.Transaction
	userId, _ := uuid.Parse(id)

	if err := tr.DB.Where("user_id = ? AND status = ?", userId, status).First(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (tr *transactionRepository) UpdateTransaction(id string, amount float64) (model.Transaction, error) {
	var transaction model.Transaction
	userId, _ := uuid.Parse(id)

    // Update grand_total di tabel transaction berdasarkan user_id
    if err := tr.DB.Model(&transaction).Where("user_id = ?", userId).Update("grand_total", amount).Error; err != nil {
        return transaction, fmt.Errorf("failed to update transaction: %v", err)
    }


	return transaction, nil
}

func (tr *transactionRepository) GetAllTransaction() ([]model.Transaction, error) {
	var transactions []model.Transaction

	if err := tr.DB.Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (tr *transactionRepository) GetAllTransactionByUserLogin(id string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := tr.DB.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
