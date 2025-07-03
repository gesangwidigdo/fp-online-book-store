package repository

import (
	"github.com/Djuanzz/go-template/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookTransactionRepository interface {
	Create(bookTransaction model.BookTransaction) (model.BookTransaction, error)
	GetByTransactionID(transactionID uuid.UUID) ([]model.BookTransaction, error)
	UpdateQuantity(transactionID uuid.UUID, bookID uuid.UUID, quantity uint64, bookPrice float64) (model.BookTransaction, error)
	Delete(transactionID uuid.UUID, bookID uuid.UUID) error
}

type bookTransactionRepository struct {
	DB *gorm.DB
}

func NewBookTransactionRepository(db *gorm.DB) BookTransactionRepository {
	return &bookTransactionRepository{
		DB: db,
	}
}

func (btr *bookTransactionRepository) Create(bookTransaction model.BookTransaction) (model.BookTransaction, error) {
	if err := btr.DB.Create(&bookTransaction).Error; err != nil {
		return model.BookTransaction{}, err
	}

	return bookTransaction, nil
}

func (btr *bookTransactionRepository) GetByTransactionID(transactionID uuid.UUID) ([]model.BookTransaction, error) {
	var bookTransactions []model.BookTransaction
	if err := btr.DB.Where("transaction_id = ?", transactionID).Find(&bookTransactions).Error; err != nil {
		return nil, err
	}

	return bookTransactions, nil
}

func (btr *bookTransactionRepository) UpdateQuantity(transactionID uuid.UUID, bookID uuid.UUID, quantity uint64, bookPrice float64) (model.BookTransaction, error) {
	var bookTransaction model.BookTransaction
	if err := btr.DB.Where("transaction_id = ? AND book_id = ?", transactionID, bookID).Take(&bookTransaction).Error; err != nil {
		return model.BookTransaction{}, err
	}

	bookTransaction.Quantity = quantity
	// Update total price
	bookTransaction.Total = bookPrice * float64(quantity)
	if err := btr.DB.Save(&bookTransaction).Error; err != nil {
		return model.BookTransaction{}, err
	}

	return bookTransaction, nil
}

func (btr *bookTransactionRepository) Delete(transactionID uuid.UUID, bookID uuid.UUID) error {
	if err := btr.DB.Where("transaction_id = ? AND book_id = ?", transactionID, bookID).Delete(&model.BookTransaction{}).Error; err != nil {
		return err
	}

	return nil
}
