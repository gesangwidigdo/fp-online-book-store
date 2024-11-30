package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/google/uuid"
)

type BookTransactionService interface {
	Create(userId string, bookTransactionReq dto.BookTransactionReq) (dto.BookTransactionRes, error)
	GetByTransactionID(transactionId uuid.UUID) ([]dto.BookTransactionRes, error)
	UpdateQuantity(transaction_id uuid.UUID, bookTransactionReq dto.BookTransactionReq) (dto.BookTransactionRes, error)
	Delete(transaction_id uuid.UUID, bookTransactionReq dto.BookTransactionDeleteReq) error
}

type bookTransactionService struct {
	bookTransactionRepo repository.BookTransactionRepository
	transactionRepo     repository.TransactionRepository
	bookRepo            repository.BookRepository
}

func NewBookTransactionService(btr repository.BookTransactionRepository, tr repository.TransactionRepository, br repository.BookRepository) BookTransactionService {
	return &bookTransactionService{
		bookTransactionRepo: btr,
		transactionRepo:     tr,
		bookRepo:            br,
	}
}

func (bts *bookTransactionService) Create(userId string, btReq dto.BookTransactionReq) (dto.BookTransactionRes, error) {
	// TRANSACTION
	// Check current transaction status by user id
	// var transaction model.Transaction
	// transaction, err := bts.transactionRepo.FindTransactionStatusByUserID(userId)
	// if err != nil {
	// 	// if record not found, create it
	// 	if err.Error() == "record not found" {
	// 		newTransaction, err := bts.transactionRepo.Create(userId)
	// 		if err != nil {
	// 			return dto.BookTransactionRes{}, err
	// 		}
	// 		transaction = newTransaction
	// 	} else {
	// 		return dto.BookTransactionRes{}, err
	// 	}
	// }

	var transaction_id uuid.UUID
	// If status is false, use current transaction id as transaction_id
	// if !transaction.Status {
	// 	transaction_id = transaction.ID
	// } else {
	// 	// If status is true, create new transaction
	// 	newTransaction, err := bts.transactionRepo.Create(userId)
	// 	if err != nil {
	// 		return dto.BookTransactionRes{}, err
	// 	}
	// 	transaction_id = newTransaction.ID
	// }

	// BOOK
	// Get book data by id
	book, err := bts.bookRepo.GetByID(btReq.BookID)
	if err != nil {
		return dto.BookTransactionRes{}, err
	}

	// Check book exist
	var emptyID [16]byte
	if book.ID == emptyID {
		return dto.BookTransactionRes{}, dto.ErrBookNotFound
	}

	// Multiply book price with quantity
	total := book.Price * float64(btReq.Quantity)

	// Create book transaction
	bookTransaction := model.BookTransaction{
		TransactionID: transaction_id,
		BookID:        btReq.BookID,
		Quantity:      btReq.Quantity,
		Total:         total,
	}

	newBookTransaction, err := bts.bookTransactionRepo.Create(bookTransaction)
	if err != nil {
		return dto.BookTransactionRes{}, err
	}

	return dto.BookTransactionRes{
		TransactionID: newBookTransaction.TransactionID,
		BookID:        newBookTransaction.BookID,
		Quantity:      newBookTransaction.Quantity,
		Total:         newBookTransaction.Total,
	}, nil
}

func (bts *bookTransactionService) GetByTransactionID(transactionId uuid.UUID) ([]dto.BookTransactionRes, error) {
	bookTransactions, err := bts.bookTransactionRepo.GetByTransactionID(transactionId)
	if err != nil {
		return nil, err
	}

	var bookTransactionList []dto.BookTransactionRes
	for _, bookTransaction := range bookTransactions {
		bookTransactionList = append(bookTransactionList, dto.BookTransactionRes{
			TransactionID: bookTransaction.TransactionID,
			BookID:        bookTransaction.BookID,
			Quantity:      bookTransaction.Quantity,
			Total:         bookTransaction.Total,
		})
	}

	return bookTransactionList, nil
}
func (bts *bookTransactionService) UpdateQuantity(transaction_id uuid.UUID, bookTransactionReq dto.BookTransactionReq) (dto.BookTransactionRes, error) {
	// Get book data
	book, err := bts.bookRepo.GetByID(bookTransactionReq.BookID)
	if err != nil {
		return dto.BookTransactionRes{}, err
	}
	
	bookTransaction, err := bts.bookTransactionRepo.UpdateQuantity(transaction_id, bookTransactionReq.BookID, bookTransactionReq.Quantity, book.Price)
	if err != nil {
		return dto.BookTransactionRes{}, err
	}

	return dto.BookTransactionRes{
		TransactionID: bookTransaction.TransactionID,
		BookID:        bookTransaction.BookID,
		Quantity:      bookTransaction.Quantity,
		Total:         bookTransaction.Total,
	}, nil
}
func (bts *bookTransactionService) Delete(transaction_id uuid.UUID, bookTransactionReq dto.BookTransactionDeleteReq) error {
	err := bts.bookTransactionRepo.Delete(transaction_id, bookTransactionReq.BookID)
	if err != nil {
		return err
	}

	return nil
}
