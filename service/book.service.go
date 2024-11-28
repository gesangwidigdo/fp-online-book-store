package service

import (
	"fmt"

	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/google/uuid"
)

type BookService interface {
	Create(bookReq dto.BookReq) (dto.BookCreateRes, error)
	GetAll() ([]dto.BookListRes, error)
	GetBySlug(slug string) (dto.BookDetailRes, error)
	Update(id uuid.UUID, bookReq dto.BookReq) (dto.BookUpdateRes, error)
	Delete(id uuid.UUID) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(br repository.BookRepository) BookService {
	return &bookService{
		bookRepo: br,
	}
}

func (bs *bookService) Create(bookReq dto.BookReq) (dto.BookCreateRes, error) {
	isISBN, err := bs.bookRepo.IsISBNExist(bookReq.ISBN)
	if err != nil {
		return dto.BookCreateRes{}, err
	}
	if isISBN {
		return dto.BookCreateRes{}, dto.ErrISBNAlreadyExists
	}

	isSlug, err := bs.bookRepo.IsSlugExist(bookReq.Slug)
	if err != nil {
		return dto.BookCreateRes{}, err
	}
	if isSlug {
		return dto.BookCreateRes{}, dto.ErrSlugAlreadyExists
	}

	book := model.Book{
		ISBN:            bookReq.ISBN,
		Title:           bookReq.Title,
		Slug:            bookReq.Slug,
		Author:          bookReq.Author,
		Summary:         bookReq.Summary,
		BookImage:       bookReq.BookImage,
		PublicationYear: bookReq.PublicationYear,
		Price:           bookReq.Price,
	}

	newBook, err := bs.bookRepo.Create(book)
	if err != nil {
		return dto.BookCreateRes{}, err
	}

	return dto.BookCreateRes{
		ID:    newBook.ID.String(),
		Title: newBook.Title,
		Price: fmt.Sprintf("%.2f", newBook.Price),
	}, nil
}

func (bs *bookService) GetAll() ([]dto.BookListRes, error) {
	books, err := bs.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var bookList []dto.BookListRes
	for _, book := range books {
		bookList = append(bookList, dto.BookListRes{
			ID:      book.ID.String(),
			Title:   book.Title,
			Author:  book.Author,
			Summary: book.Summary,
			Price:   book.Price,
		})
	}

	return bookList, nil
}

func (bs *bookService) GetBySlug(slug string) (dto.BookDetailRes, error) {
	book, err := bs.bookRepo.GetBySlug(slug)
	if err != nil {
		return dto.BookDetailRes{}, err
	}

	return dto.BookDetailRes{
		ID:              book.ID.String(),
		ISBN:            book.ISBN,
		Title:           book.Title,
		Author:          book.Author,
		Summary:         book.Summary,
		BookImage:       book.BookImage,
		PublicationYear: book.PublicationYear,
		Price:           book.Price,
	}, nil
}

func (bs *bookService) Update(id uuid.UUID, bookReq dto.BookReq) (dto.BookUpdateRes, error) {
	// Get the existing book data
	existingBook, err := bs.bookRepo.GetByID(id)
	if err != nil {
		return dto.BookUpdateRes{}, err
	}

	// If exisiting book's isbn is not equal to the new book's isbn,
	// check if the new book's isbn already exists
	if existingBook.ISBN != bookReq.ISBN {
		isISBN, err := bs.bookRepo.IsISBNExist(bookReq.ISBN)
		if err != nil {
			return dto.BookUpdateRes{}, err
		}
		if isISBN {
			return dto.BookUpdateRes{}, dto.ErrISBNAlreadyExists
		}

		isSlug, err := bs.bookRepo.IsSlugExist(bookReq.Slug)
		if err != nil {
			return dto.BookUpdateRes{}, err
		}
		if isSlug {
			return dto.BookUpdateRes{}, dto.ErrSlugAlreadyExists
		}
	}

	book := model.Book{
		ISBN:            bookReq.ISBN,
		Title:           bookReq.Title,
		Slug:            bookReq.Slug,
		Author:          bookReq.Author,
		Summary:         bookReq.Summary,
		BookImage:       bookReq.BookImage,
		PublicationYear: bookReq.PublicationYear,
		Price:           bookReq.Price,
	}

	updatedBook, err := bs.bookRepo.Update(id, book)
	if err != nil {
		return dto.BookUpdateRes{}, err
	}

	return dto.BookUpdateRes{
		ID:              id.String(),
		ISBN:            updatedBook.ISBN,
		Title:           updatedBook.Title,
		Author:          updatedBook.Author,
		Summary:         updatedBook.Summary,
		BookImage:       updatedBook.BookImage,
		PublicationYear: updatedBook.PublicationYear,
		Price:           updatedBook.Price,
	}, nil
}

func (bs *bookService) Delete(id uuid.UUID) error {
	err := bs.bookRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
