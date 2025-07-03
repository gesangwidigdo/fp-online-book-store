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
	GetTransactionWithBooksByID(id string) (model.Transaction, error)
	UpdateTransaction(id string, amount float64) (model.Transaction, error)
	UpdateStatus(id string, status string) (model.Transaction, error)
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

	if err := tr.DB.Where("user_id = ? AND status = ?", userId, status).
		Find(&transaction).Error; err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (tr *transactionRepository) UpdateTransaction(id string, amount float64) (model.Transaction, error) {
	var transaction model.Transaction

	// Update grand_total di tabel transaction berdasarkan user_id
	if err := tr.DB.Model(&transaction).Where("id = ?", id).Update("grand_total", amount).Error; err != nil {
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

	if err := tr.DB.Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&transactions).Error; err != nil {
			return nil, err
	}

	return transactions, nil
}

func (tr *transactionRepository) GetTransactionWithBooksByID(id string) (model.Transaction, error) {
	var transaction model.Transaction

	if err := tr.DB.Preload("BookTransaction.Book").Where("id = ?", id).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (tr *transactionRepository) UpdateStatus(id string, status string) (model.Transaction, error) {
    var transaction model.Transaction

    // Cari transaksi berdasarkan ID
    if err := tr.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
        return transaction, err
    }

    // Perbarui status transaksi
    if err := tr.DB.Model(&transaction).Update("status", status).Error; err != nil {
        return transaction, err
    }

    return transaction, nil
}