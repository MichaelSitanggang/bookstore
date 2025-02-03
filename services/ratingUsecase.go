package services

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
)

type RatingServices interface {
	TambangRating(rating entities.Rating) error
	UbahBookRating(bookID int) error
}

type ratingServices struct {
	ratingRepo repositories.RatingRepo
}

func NewServicesRating(ratingRepo repositories.RatingRepo) RatingServices {
	return &ratingServices{ratingRepo: ratingRepo}
}

func (s *ratingServices) TambangRating(rating entities.Rating) error {
	return s.ratingRepo.AddRating(rating)
}

func (s *ratingServices) UbahBookRating(bookID int) error {
	averageRating, reviewCount, err := s.ratingRepo.GetTotalRatingByBook(bookID)
	if err != nil {
		return err
	}
	return s.ratingRepo.UpdateBookRating(bookID, averageRating, reviewCount)
}
