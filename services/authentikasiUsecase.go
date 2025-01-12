package services

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/MichaelSitanggang/bookstore/entities"
	"github.com/MichaelSitanggang/bookstore/repositories"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type AuthService interface {
	Login(email, password string) (*entities.User, error)
	Register(namalengkap string, email string, umur int, password string) (string, error)
	VerifikasiOtp(otp string) error
	SendOtp(email string) (string, error)
	LoginAdmin(email, password string) (*entities.Admin, error)
}

type authService struct {
	repo repositories.AuthRepo
}

func NewAuthService(repo repositories.AuthRepo) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SendOtp(email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(999999-100000) + 100000
	otpStr := strconv.Itoa(otp)

	user, _ := s.repo.FindByEmail(email)
	user.OTP = otpStr
	if err := s.repo.UpdateOtp(user); err != nil {
		return "", err
	}

	from := mail.NewEmail("BookStore", "michaelsitanggang37@gmail.com")
	subject := "Verifikasi Otp"
	to := mail.NewEmail("User", email)
	plainTextContent := "Code Otp"
	htmlContent := "<strong>Code Otp Kamu adalah " + otpStr + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return "", err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return otpStr, nil
	}
	return "", errors.New("otp tidak ada")
}

func (s *authService) Register(namalengkap string, email string, umur int, password string) (string, error) {
	emails, _ := s.repo.FindByEmail(email)
	if emails != nil {
		return "", errors.New("email sudah terdaftar")
	}
	user := entities.User{
		NamaLengkap: namalengkap,
		Email:       email,
		Umur:        umur,
		Password:    password,
	}
	if len(user.Password) < 7 {
		return "", errors.New("password harus lebih dari 7 karakter")
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return "", err
	}
	sendOtp, err := s.SendOtp(user.Email)
	if err != nil {
		return "", err
	}
	user.OTP = sendOtp
	if err := s.repo.UpdateOtp(&user); err != nil {
		return "", err
	}
	return sendOtp, nil
}

func (s *authService) VerifikasiOtp(otp string) error {
	otps, _ := s.repo.FindByOtp(otp)
	if otp == "" || otp != otps.OTP {
		return errors.New("otp tidak benar")
	}
	otps.StatusOtp = true
	if err := s.repo.UpdateOtp(otps); err != nil {
		return err
	}
	return nil
}

func (s *authService) Login(email, password string) (*entities.User, error) {
	user, _ := s.repo.FindByEmail(email)
	if password != user.Password {
		return nil, errors.New("password salah")
	}
	if !user.StatusOtp {
		return nil, errors.New("otp belum diverifikasi")
	}
	return user, nil
}

func (s *authService) LoginAdmin(email, password string) (*entities.Admin, error) {
	admin, _ := s.repo.FindByEmailAdmin(email)
	if password != admin.Password {
		return nil, errors.New("password salah")
	}
	return admin, nil
}
