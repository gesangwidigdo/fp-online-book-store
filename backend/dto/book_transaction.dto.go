package dto

import (
	"errors"

	"github.com/google/uuid"
)

type BookTransactionReq struct {
	BookID        uuid.UUID `json:"book_id" form:"book_id" binding:"required"`
	Quantity      uint64    `json:"quantity" form:"quantity" binding:"required"`
}

type BookTransactionDeleteReq struct {
	BookID        uuid.UUID `json:"book_id" form:"book_id" binding:"required"`
}

type BookTransactionRes struct {
	BookID        uuid.UUID `json:"book_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Quantity      uint64    `json:"quantity"`
	Total         float64   `json:"total"`
}

const (
	MSG_BOOK_TRANSACTION_CREATE_SUCCESS  = "book transaction created successfully"
	MSG_BOOK_TRANSACTION_UPDATE_SUCCESS  = "book transaction updated successfully"
	MSG_BOOK_TRANSACTION_DELETE_SUCCESS  = "book transaction deleted successfully"
	MSG_BOOK_TRANSACTION_GET_ALL_SUCCESS = "get all book transactions by transaction id successfully"

	MSG_BOOK_TRANSACTION_CREATE_FAILED  = "book transaction failed to create"
	MSG_BOOK_TRANSACTION_UPDATE_FAILED  = "book transaction failed to update"
	MSG_BOOK_TRANSACTION_DELETE_FAILED  = "book transaction failed to delete"
	MSG_BOOK_TRANSACTION_GET_ALL_FAILED = "failed to get all book transactions by transaction id"
)

var (
	ErrBookTransactionNotFound = errors.New("book transaction not found")
	ErrBookTransactionExist    = errors.New("book transaction already exist")
)
