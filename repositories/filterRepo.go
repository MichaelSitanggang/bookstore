package repositories

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/gorm"
)

type FilterRepo interface {
	FilterBooks(judul string, tahun int) (entities.Book, error)
	FilterByPenjualan(limit int) ([]entities.Book, error)
}

type filterRepo struct {
	db *gorm.DB
}

func NewFilterBook(db *gorm.DB) FilterRepo {
	return &filterRepo{db: db}
}

func (r *filterRepo) FilterBooks(judul string, tahun int) (entities.Book, error) {
	var book entities.Book
	query := r.db.Model(entities.Book{})
	if judul != "" {
		query = query.Where("judul LIKE ?", "%"+judul+"%")
	}
	if tahun > 0 {
		query = query.Where("year = ?", tahun)
	}
	err := query.Find(&book).Error
	return book, err
}

func (r *filterRepo) FilterByPenjualan(limit int) ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Order("Penjualan DESC").Limit(limit).Find(&books).Error
	return books, err
}
