package repositories

import (
	"fiber-crud-auth/databases"
	"fiber-crud-auth/models"
)

type BookRepository interface {
    Create(book *models.Book) error
    FindAll() ([]models.Book, error)
    FindByID(id uint) (*models.Book, error)
    Update(book *models.Book) error
    Delete(id uint) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
    return &bookRepository{}
}

func (r *bookRepository) Create(book *models.Book) error {
    return databases.DB.Create(book).Error
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
    var books []models.Book
    err := databases.DB.Find(&books).Error
    return books, err
}

func (r *bookRepository) FindByID(id uint) (*models.Book, error) {
    var book models.Book
    err := databases.DB.First(&book, id).Error
    return &book, err
}

func (r *bookRepository) Update(book *models.Book) error {
    return databases.DB.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
    return databases.DB.Delete(&models.Book{}, id).Error
}
