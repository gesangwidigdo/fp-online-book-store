package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/repository"
)

type TransactionService interface {
	CreateTransaction(id string) (dto.TransactionCreateRes, error)
	GetTransactionStatus(id string) (dto.TransactionStatusRes, error)
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
	trans, err := ts.transactionRepo.FindTransactionStatusByUserID(id)
	if err != nil {
		return dto.TransactionStatusRes{}, err
	}

	return dto.TransactionStatusRes{
		GrandTotal: trans.GrandTotal,
		CreatedAt: trans.CreatedAt,
		Status: trans.Status,
	}, nil
}