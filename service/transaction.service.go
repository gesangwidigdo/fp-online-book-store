package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/repository"
)

type TransactionService interface {
	CreateTransaction(id string) (dto.TransactionCreateRes, error)
	GetTransactionStatus(id string) (dto.TransactionStatusRes, error)
	GetAllTransactionByUserLogin(id string) (dto.TransactionListRes, error)
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
		CreatedAt: trans.CreatedAt,
	}, nil
}

func (ts *transactionService) GetTransactionStatus(id string) (dto.TransactionStatusRes, error) {
	trans, err := ts.transactionRepo.FindTransactionStatusByUserID(id, "draft")
	if err != nil {
		return dto.TransactionStatusRes{}, err
	}

	return dto.TransactionStatusRes{
		GrandTotal: trans.GrandTotal,
		CreatedAt: trans.CreatedAt,
		Status: trans.Status,
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
			GrandTotal: t.GrandTotal,
			CreatedAt: t.CreatedAt,
			Status: t.Status,
		})
	}

	return dto.TransactionListRes{
		Transactions: transactionStatusRes,
	}, nil
}


func (ts *transactionService) UpdateTransaction() error{
	return nil
}