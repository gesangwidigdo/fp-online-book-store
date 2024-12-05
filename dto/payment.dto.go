package dto

type CreatePaymentRequest struct {
	Amount int64 `json:"amount" form:"amount" binding:"required"`
	TransactionID string `json:"transaction_id" form:"transaction_id"`
}

type CreatePaymentResponse struct {
	Amount int64 `json:"amount"`
	Status string `json:"status"`
	TransactionID string `json:"transaction_id"`
}

type PaymentNotificationResponse struct {
	Transaction_time string `json:"transaction_time"`
	Transaction_status string `json:"transaction_status"`
	Transaction_id string `json:"transaction_id"`
	Status_message string `json:"status_message"`
	Status_code string `json:"status_code"`
	Signature_key string `json:"signature_key"`
	Settlement_time string `json:"settlement_time"`
	Payment_type string `json:"payment_type"`
	Order_id string `json:"order_id"`
	Merchant_id string `json:"merchant_id"`
	Gross_amount string `json:"gross_amount"`
	Fraud_status string `json:"fraud_status"`
	Currency string `json:"currency"`
}

type PaymentPayload struct {
	Gross_amount int64 `json:"gross_amount"`
	Order_id string `json:"order_id"`
}

const (
	MSG_CREATE_PAYMENT_SUCCESS        = "create payment success"

	MSG_CREATE_PAYMENT_FAILED = "create payment failed"

)