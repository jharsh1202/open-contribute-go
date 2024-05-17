package repositories

import (
	"open-contribute/pkg/db/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := r.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) UpdateBook(book *models.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepository) DeleteBook(book *models.Book) error {
	return r.DB.Delete(book).Error
}
