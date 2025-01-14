package services

import (
	"errors"

	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
)

type BookService interface {
	GetAllBook() ([]entities.Book, error)
	GetBookById(id int) (*entities.Book, error)
	CreateBook(gambar string, judul string, author string, year int, harga float64, stok int) error
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
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("data kosong")
	}
	return book, nil
}

func (s *bookService) CreateBook(gambar string, judul string, author string, year int, harga float64, stok int) error {
	book := entities.Book{
		Gambar: gambar,
		Judul:  judul,
		Author: author,
		Year:   year,
		Harga:  harga,
		Stok:   stok,
	}
	return s.repo.CreateBook(&book)
}
