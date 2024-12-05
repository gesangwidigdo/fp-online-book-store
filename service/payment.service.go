package service

import (
	"fmt"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/google/uuid"
)

type PaymentService interface {
	CreateStandard(paymentReq dto.CreatePaymentRequest, id string) (dto.CreatePaymentResponse, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
	transRepo repository.TransactionRepository
}

func NewPaymentService(pr repository.PaymentRepository, tr repository.TransactionRepository) PaymentService {
	return &paymentService{
		paymentRepo: pr,
		transRepo: tr,
	}
}

// --- JADI ANGGEPANNYA INI TU BAKAL BAYAR TRANSAKSINE TRUS LGSG SUCCESS
func (ps *paymentService) CreateStandard(paymentReq dto.CreatePaymentRequest, userId string) (dto.CreatePaymentResponse, error) {
	// --- BUAT DAPETIN TRANSACTION YANG STATUSNYA DRAFT DAN ID NYA
	trans, err := ps.transRepo.FindTransactionStatusByUserID(userId, "draft")
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	transId, _ := uuid.Parse(trans.ID.String())

	// transactions, err := ps.transactionRepo.GetTransactionWithBooksByID(transIdStr)
	transaction, err := ps.transRepo.GetTransactionWithBooksByID(transId.String())
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	// --- ANGGEPANE CREATE LGSG SUCCESS
	payment := model.Payment{
		TransactionID: transId,
		Amount: int64(transaction.GrandTotal),
		Status: "success",
	}

	res, err := ps.paymentRepo.Create(payment)
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}

	transIdStr := res.TransactionID.String()
	fmt.Println(transIdStr)

	// --- KARNA PAYMEN E DAH SUCCESS, BRATI LANJUT KE UPDATE STATUS TRANSACTION KE SUCCESS
	transRes, err := ps.transRepo.UpdateStatus(transIdStr, "success")
	if err != nil {
		return dto.CreatePaymentResponse{}, err
	}
	
	transIdStr = transRes.ID.String()
	fmt.Println(transIdStr)
	
	return dto.CreatePaymentResponse{
		TransactionID: transIdStr,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
	
}