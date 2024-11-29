package repository

import (
	"github.com/Djuanzz/go-template/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book model.Book) (model.Book, error)
	GetAll() ([]model.Book, error)
	GetBySlug(slug string) (model.Book, error)
	GetByID(id uuid.UUID) (model.Book, error)
	Update(id uuid.UUID, book model.Book) (model.Book, error)
	Delete(id uuid.UUID) error

	IsISBNExist(isbn string) (bool, error)
	IsSlugExist(slug string) (bool, error)
}

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}

func (br *bookRepository) Create(book model.Book) (model.Book, error) {
	if err := br.DB.Create(&book).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (br *bookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	if err := br.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br *bookRepository) GetBySlug(slug string) (model.Book, error) {
	var book model.Book
	if err := br.DB.Where("slug = ?", slug).Take(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Book{}, nil
		}
		return model.Book{}, err
	}

	return book, nil
}

func (br *bookRepository) GetByID(id uuid.UUID) (model.Book, error) {
	var book model.Book
	if err := br.DB.Where("id = ?", id).Take(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Book{}, nil
		}
		return model.Book{}, err
	}

	return book, nil
}

func (br *bookRepository) Update(id uuid.UUID, book model.Book) (model.Book, error) {
	var existingBook model.Book
	if err := br.DB.Where("id = ?", id).Take(&existingBook).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Book{}, nil
		}
		return model.Book{}, err
	}

	if err := br.DB.Model(&existingBook).Updates(book).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (br *bookRepository) Delete(id uuid.UUID) error {
	if err := br.DB.Where("id = ?", id).Delete(&model.Book{}).Error; err != nil {
		return err
	}

	return nil
}

func (br *bookRepository) IsISBNExist(isbn string) (bool, error) {
	var book model.Book
	if err := br.DB.Where("isbn = ?", isbn).Take(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (br *bookRepository) IsSlugExist(slug string) (bool, error) {
	var book model.Book
	if err := br.DB.Where("slug = ?", slug).Take(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}