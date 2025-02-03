package services

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
)

type RatingServices interface {
	TambahReview(UserID int, BookID int, Rating int, ulasan string) error
	UpdateRating(bookID int) error
}

type ratingServices struct {
	ratingRepo repositories.RatingRepo
}

func NewServicesRating(ratingRepo repositories.RatingRepo) RatingServices {
	return &ratingServices{ratingRepo: ratingRepo}
}

func (s *ratingServices) TambahReview(UserID int, BookID int, Rating int, ulasan string) error {
	rating := entities.Rating{
		UserID: UserID,
		BookID: BookID,
		Rating: Rating,
		Ulasan: ulasan,
	}
	if err := s.ratingRepo.AddRating(rating); err != nil {
		return err
	}
	return s.UpdateRating(BookID)
}

func (s *ratingServices) UpdateRating(bookID int) error {
	totalReviews, err := s.ratingRepo.CountReviewByBook(bookID)
	if err != nil {
		return err
	}
	totalRating, err := s.ratingRepo.SumReviewByBook(bookID)
	if err != nil {
		return err
	}
	newRating := 0.0
	if totalReviews > 0 {
		newRating = totalRating / float64(totalReviews)
	}
	err = s.ratingRepo.UpdateBookRating(bookID, newRating)
	if err != nil {
		return err
	}
	return err
}
