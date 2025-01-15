package services

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
)

type FilterBookService interface {
	CariBooks(judul string, tahun int) (entities.Book, error)
}

type filterBookService struct {
	repo repositories.FilterRepo
}

func NewFilterService(repo repositories.FilterRepo) FilterBookService {
	return &filterBookService{repo: repo}
}

func (s *filterBookService) CariBooks(judul string, tahun int) (entities.Book, error) {
	return s.repo.FilterBooks(judul, tahun)
}
