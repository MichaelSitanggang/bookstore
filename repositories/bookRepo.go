package repositories

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/gorm"
)

type BookRepo interface {
	FindAll() ([]entities.Book, error)
	FindByID(id int) (*entities.Book, error)
	CreateBook(book *entities.Book) error
}

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{db: db}
}

func (r *bookRepo) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	r.db.Preload("Reviews.User").Find(&books)
	return books, nil
}

func (r *bookRepo) FindByID(id int) (*entities.Book, error) {
	var books entities.Book
	if err := r.db.Preload("Reviews.User").Where("id = ?", id).First(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}

func (r *bookRepo) CreateBook(book *entities.Book) error {
	return r.db.Create(book).Error
}
