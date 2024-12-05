package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/repository"
)

type TransactionService interface {
	CreateTransaction(id string) (dto.TransactionCreateRes, error)
	GetTransactionStatus(id string) (dto.TransactionStatusRes, error)
	GetAllTransactionByUserLogin(id string) (dto.TransactionListRes, error)
	GetTransactionWithBooksByID(id string) (dto.TransactionWithBooksRes, error)
	CalculateGrandTotal(id string) (dto.TransactionCalculateGrandTotalRes, error)
	GetTransactionWithBooksByUserLogin(userId string) (dto.TransactionWithBooksRes, error)
	CalculateGrandTotalByUserLogin(userId string) (dto.TransactionCalculateGrandTotalRes, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(tr repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: tr,
	}
}

func (ts *transactionService) CreateTransaction(id string) (dto.TransactionCreateRes, error) {
	availTrans, _ := ts.transactionRepo.FindTransactionStatusByUserID(id, "draft")

	if availTrans.Status == "draft" {
		return dto.TransactionCreateRes{}, dto.ERR_TRANSACTION_ALREADY_EXISTS
	}

	trans, err := ts.transactionRepo.Create(id)

	if err != nil {
		return dto.TransactionCreateRes{}, err
	}

	return dto.TransactionCreateRes{
		GrandTotal: trans.GrandTotal,
		CreatedAt:  trans.CreatedAt,
	}, nil
}

func (ts *transactionService) GetTransactionStatus(id string) (dto.TransactionStatusRes, error) {
	trans, err := ts.transactionRepo.FindTransactionStatusByUserID(id, "draft")
	if err != nil {
		return dto.TransactionStatusRes{}, err
	}

	return dto.TransactionStatusRes{
		TransId: trans.ID.String(),
		GrandTotal: trans.GrandTotal,
		CreatedAt:  trans.CreatedAt,
		Status:     trans.Status,
	}, nil
}

func (ts *transactionService) GetAllTransactionByUserLogin(id string) (dto.TransactionListRes, error) {
	transactions, err := ts.transactionRepo.GetAllTransactionByUserLogin(id)
	if err != nil {
		return dto.TransactionListRes{}, err
	}

	var transactionStatusRes []dto.TransactionStatusRes
	for _, t := range transactions {
		transactionStatusRes = append(transactionStatusRes, dto.TransactionStatusRes{
			TransId: t.ID.String(),
			GrandTotal: t.GrandTotal,
			CreatedAt:  t.CreatedAt,
			Status:     t.Status,
		})
	}

	return dto.TransactionListRes{
		Transactions: transactionStatusRes,
	}, nil
}

func (ts *transactionService) UpdateTransaction() error {
	return nil
}

func (ts *transactionService) GetTransactionWithBooksByID(id string) (dto.TransactionWithBooksRes, error) {
	transactions, err := ts.transactionRepo.GetTransactionWithBooksByID(id)
	if err != nil {
		return dto.TransactionWithBooksRes{}, err
	}

	return dto.TransactionWithBooksRes{
		ID:         transactions.ID.String(),
		GrandTotal: transactions.GrandTotal,
		BookList: func() []dto.BookToTransactionRes {
			var bookList []dto.BookToTransactionRes
			for _, b := range transactions.BookTransaction {
				bookList = append(bookList, dto.BookToTransactionRes{
					ID:        b.BookID,
					Title:     b.Book.Title,
					BookImage: b.Book.BookImage,
					Price:     b.Book.Price,
					Quantity:  b.Quantity,
					Total:     b.Total,
				})
			}
			return bookList
		}(),
	}, nil
}

func (ts *transactionService) GetTransactionWithBooksByUserLogin(userId string) (dto.TransactionWithBooksRes, error) {
	trans, err := ts.transactionRepo.FindTransactionStatusByUserID(userId, "draft")
	if err != nil {
		return dto.TransactionWithBooksRes{}, err
	}

	transIdStr := trans.ID.String()
	
	transactions, err := ts.transactionRepo.GetTransactionWithBooksByID(transIdStr)
	if err != nil {
		return dto.TransactionWithBooksRes{}, err
	}

	return dto.TransactionWithBooksRes{
		ID:         transactions.ID.String(),
		GrandTotal: transactions.GrandTotal,
		BookList: func() []dto.BookToTransactionRes {
			var bookList []dto.BookToTransactionRes
			for _, b := range transactions.BookTransaction {
				bookList = append(bookList, dto.BookToTransactionRes{
					ID:        b.BookID,
					Title:     b.Book.Title,
					BookImage: b.Book.BookImage,
					Price:     b.Book.Price,
					Quantity:  b.Quantity,
					Total:     b.Total,
				})
			}
			return bookList
		}(),
	}, nil
}

func (ts *transactionService) CalculateGrandTotal(id string) (dto.TransactionCalculateGrandTotalRes, error) {
	transactions, err := ts.transactionRepo.GetTransactionWithBooksByID(id)
	if err != nil {
		return dto.TransactionCalculateGrandTotalRes{}, err
	}

	var grandTotal float64
	for _, bt := range transactions.BookTransaction {
		grandTotal += bt.Total
	}

	updatedGrandTotal, err := ts.transactionRepo.UpdateTransaction(id, grandTotal)
	if err != nil {
		return dto.TransactionCalculateGrandTotalRes{}, err
	}

	return dto.TransactionCalculateGrandTotalRes{
		ID:         transactions.ID.String(),
		GrandTotal: updatedGrandTotal.GrandTotal,
	}, nil
}

func (ts *transactionService) CalculateGrandTotalByUserLogin(userId string) (dto.TransactionCalculateGrandTotalRes, error) {
	trans, err := ts.transactionRepo.FindTransactionStatusByUserID(userId, "draft")
	if err != nil {
		return dto.TransactionCalculateGrandTotalRes{}, err
	}

	transIdStr := trans.ID.String()
	
	transactions, err := ts.transactionRepo.GetTransactionWithBooksByID(transIdStr)
	if err != nil {
		return dto.TransactionCalculateGrandTotalRes{}, err
	}

	var grandTotal float64
	for _, bt := range transactions.BookTransaction {
		grandTotal += bt.Total
	}

	updatedGrandTotal, err := ts.transactionRepo.UpdateTransaction(transIdStr, grandTotal)
	if err != nil {
		return dto.TransactionCalculateGrandTotalRes{}, err
	}

	return dto.TransactionCalculateGrandTotalRes{
		ID:         transactions.ID.String(),
		GrandTotal: updatedGrandTotal.GrandTotal,
	}, nil
}

