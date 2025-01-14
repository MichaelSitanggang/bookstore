package services

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
)

type BookService interface {
	GetAllBook() ([]entities.Book, error)
	GetBookById(id int) (*entities.Book, error)
	CreateBook(book entities.Book) error
}

type bookService struct {
	repo repositories.BookRepo
}

func NewBookService(repo repositories.BookRepo) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBook() ([]entities.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookById(id int) (*entities.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book entities.Book) error {
	return s.repo.CreateBook(&book)
}
