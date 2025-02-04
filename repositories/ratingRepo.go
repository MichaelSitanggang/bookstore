package repositories

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/gorm"
)

type RatingRepo interface {
	AddRating(rating entities.Rating) error
	UpdateBookRating(BookID int, RatingBaru float64) error
	SumReviewByBook(BookID int) (float64, error)
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) RatingRepo {
	return &ratingRepo{db: db}
}

func (r *ratingRepo) AddRating(rating entities.Rating) error {
	return r.db.Create(&rating).Error
}

func (r *ratingRepo) UpdateBookRating(BookID int, RatingBaru float64) error {
	var book entities.Book
	if err := r.db.First(&book, BookID).Error; err != nil {
		return err
	}
	book.Review = RatingBaru
	return r.db.Save(&book).Error
}
func (r *ratingRepo) SumReviewByBook(BookID int) (float64, error) {
	var totalRating float64
	err := r.db.Model(&entities.Rating{}).Where("book_id = ?", BookID).Select("SUM(rating)").Scan(&totalRating).Error
	if err != nil {
		return 0, err
	}
	return totalRating, nil
}
