package services

import (
	"fiber-crud-auth/models"
	"fiber-crud-auth/repositories"
)

type BookService interface {
    CreateBook(book *models.Book) error
    GetAllBooks() ([]models.Book, error)
    GetBookByID(id uint) (*models.Book, error)
    UpdateBook(book *models.Book) error
    DeleteBook(id uint) error
}

type bookService struct {
    repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
    return &bookService{repo}
}

func (s *bookService) CreateBook(book *models.Book)error{
	return s.repo.Create(book)
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
    return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
    return s.repo.FindByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
    return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
    return s.repo.Delete(id)
}





