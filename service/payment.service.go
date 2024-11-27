package service

import "github.com/Djuanzz/go-template/repository"

type PaymentService interface {

}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(pr repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepo: pr,
	}
}	