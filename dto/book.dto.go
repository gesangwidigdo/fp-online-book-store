package dto

import (
	"errors"

	"github.com/google/uuid"
)

type BookReq struct {
	ISBN            string  `json:"isbn" form:"isbn" binding:"required"`
	Title           string  `json:"title" form:"title" binding:"required"`
	Slug            string  `json:"slug" form:"slug" binding:"required"`
	Author          string  `json:"author" form:"author" binding:"required"`
	Summary         string  `json:"summary" form:"summary" binding:"required"`
	BookImage       string  `json:"book_image" form:"book_image" binding:"required"`
	PublicationYear uint64  `json:"publication_year" form:"publication_year" binding:"required"`
	Price           float64 `json:"price" form:"price" binding:"required"`
}

type BookCreateRes struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Price string    `json:"price"`
}

type BookListRes struct {
	ID        uuid.UUID `json:"id"`
	BookImage string    `json:"book_image"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Author    string    `json:"author"`
	Summary   string    `json:"summary"`
	Price     float64   `json:"price"`
}

type BookDetailRes struct {
	ID              uuid.UUID `json:"id"`
	ISBN            string    `json:"isbn"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Summary         string    `json:"summary"`
	BookImage       string    `json:"book_image"`
	PublicationYear uint64    `json:"publication_year"`
	Price           float64   `json:"price"`
}

type BookUpdateRes struct {
	ID              uuid.UUID `json:"id"`
	ISBN            string    `json:"isbn"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Summary         string    `json:"summary"`
	BookImage       string    `json:"book_image"`
	PublicationYear uint64    `json:"publication_year"`
	Price           float64   `json:"price"`
}

type BookToTransactionRes struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	BookImage string    `json:"book_image"`
	Price     float64   `json:"price"`
	Quantity  uint64    `json:"quantity"`
	Total     float64   `json:"total"`
}

const (
	MSG_BOOK_CREATE_SUCCESS      = "book created successfully"
	MSG_BOOK_UPDATE_SUCCESS      = "book updated successfully"
	MSG_BOOK_DELETE_SUCCESS      = "book deleted successfully"
	MSG_BOOK_GET_ALL_SUCCESS     = "get all books successfully"
	MSG_BOOK_GET_BY_SLUG_SUCCESS = "get book by slug successfully"

	MSG_BOOK_CREATE_FAILED      = "book creation failed"
	MSG_BOOK_UPDATE_FAILED      = "book update failed"
	MSG_BOOK_DELETE_FAILED      = "book delete failed"
	MSG_BOOK_GET_ALL_FAILED     = "get all books failed"
	MSG_BOOK_GET_BY_SLUG_FAILED = "get book by slug failed"

	MSG_BOOK_NOT_FOUND = "book not found"
)

var (
	ErrBookNotFound      = errors.New("book not found")
	ErrISBNAlreadyExists = errors.New("book with specified ISBN already exists")
	ErrSlugAlreadyExists = errors.New("book with specified slug already exists")
)
