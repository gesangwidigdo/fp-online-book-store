package dto

import "time"

type TransactionReq struct {

}

type TransactionCreateRequest struct {

}

type TransactionCreateRes struct {
	GrandTotal float64 `json:"grand_total"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionStatusRes struct {
	GrandTotal float64 `json:"grand_total"`
	CreatedAt time.Time `json:"created_at"`
	Status bool `json:"status"`
}

const (
	MSG_TRANSACTION_CREATE_SUCCESS = "transaction created successfully"
	MSG_TRANSACTION_UPDATE_SUCCESS = "transaction updated successfully"
	MSG_TRANSACTION_DELETE_SUCCESS = "transaction deleted successfully"
	MSG_TRANSACTION_STATUS_SUCCESS = "transaction status retrieved successfully"

	MSG_TRANSACTION_CREATE_FAILED = "transaction failed to create"
	MSG_TRANSACTION_UPDATE_FAILED = "transaction failed to update"
	MSG_TRANSACTION_STATUS_FAILED = "transaction status failed to get"
)