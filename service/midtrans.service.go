package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	GenerateSnapUrl(paymentReq dto.CreatePaymentRequest) (string, error)
}

type midtransService struct {
	client *snap.Client
}

func NewMidtransService(client *snap.Client) MidtransService {
	return &midtransService{
		client: client,
	}
}

func (ms *midtransService) GenerateSnapUrl(paymentReq dto.CreatePaymentRequest) (string, error) {
	req := & snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.New().String(), 
			GrossAmt: int64(paymentReq.Amount),
		}, 
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapRes, err := ms.client.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	// --- TESTING
	return snapRes.RedirectURL, nil
}