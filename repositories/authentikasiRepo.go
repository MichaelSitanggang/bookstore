package repositories

import (
	"errors"

	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/gorm"
)

type AuthRepo interface {
	FindByEmail(email string) (*entities.User, error)
	UpdateOtp(user *entities.User) error
	CreateUser(user *entities.User) error
	FindByOtp(otp string) (*entities.User, error)
	FindByEmailAdmin(email string) (*entities.Admin, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) FindByEmailAdmin(email string) (*entities.Admin, error) {
	var admin entities.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil

}

func (r *authRepo) UpdateOtp(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *authRepo) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *authRepo) FindByOtp(otp string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("otp = ?", otp).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
