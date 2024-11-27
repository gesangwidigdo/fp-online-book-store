package dto

type CreatePaymentRequest struct {
	Change float64 `json:"change" form:"change" binding:"required"`
}

const (
	MSG_CREATE_PAYMENT_SUCCESS        = "create payment success"

	MSG_CREATE_PAYMENT_FAILED = "create payment failed"

)