package repositories

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/gorm"
)

type RatingRepo interface {
	AddRating(rating entities.Rating) error
	GetTotalRatingByBook(bookID int) (float64, int, error)
	UpdateBookRating(bookID int, totalRating float64, reviewCount int) error
}

type ratingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) RatingRepo {
	return &ratingRepo{db: db}
}

func (r *ratingRepo) AddRating(rating entities.Rating) error {
	err := r.db.Create(rating).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ratingRepo) GetTotalRatingByBook(bookID int) (float64, int, error) {
	var rating []entities.Rating
	var totalRating int

	err := r.db.Where("book_id = ?", bookID).Find(&rating).Error
	if err != nil {
		return 0, 0, err
	}
	for _, r := range rating {
		totalRating += r.Rating
	}
	count := len(rating)
	if count == 0 {
		return 0, 0, nil
	}
	averangeRating := float64(totalRating) / float64(count)
	return averangeRating, count, nil
}

func (r *ratingRepo) UpdateBookRating(bookID int, totalRating float64, reviewCount int) error {
	return r.db.Model(&entities.Book{}).Where("id = ?", bookID).Updates(map[string]interface{}{
		"total_rating": totalRating,
		"review":       reviewCount,
	}).Error
}
