package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/bookstore/middlewares"
	"github.com/MichaelSitanggang/bookstore/services"
	"github.com/gin-gonic/gin"
)

type AuthControl struct {
	usecase services.AuthService
}

func NewAuthControl(usecase services.AuthService) *AuthControl {
	return &AuthControl{usecase: usecase}
}

func (ctrl *AuthControl) Registers(c *gin.Context) {
	type input struct {
		NamaLengkap string `json:"nama_lengkap"`
		Email       string `json:"email"`
		Umur        int    `json:"umur"`
		Password    string `json:"password"`
	}
	var inputan input
	if err := c.ShouldBindJSON(&inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	_, err := ctrl.usecase.Register(inputan.NamaLengkap, inputan.Email, inputan.Umur, inputan.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil registrasi"})
}
func (ctrl *AuthControl) VerifikasiOtps(c *gin.Context) {
	var inputan struct {
		Otp string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	if err := ctrl.usecase.VerifikasiOtp(inputan.Otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Verifikasi Otp Berhasil lanjut login"})
}

func (ctrl *AuthControl) Login(c *gin.Context) {
	type input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var inputan input
	if err := c.ShouldBindJSON(&inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	user, err := ctrl.usecase.Login(inputan.Email, inputan.Password)
	if err == nil && user != nil {
		tokenUser, err := middlewares.GenerateJwt(user.ID, "user")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "login berhasil", "token": tokenUser})
		return
	}

	admin, erradmin := ctrl.usecase.LoginAdmin(inputan.Email, inputan.Password)
	if erradmin == nil && admin != nil {
		tokenAdmin, err := middlewares.GenerateJwt(admin.ID, "admin")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "berhasil login", "token": tokenAdmin})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Login gagal"})
}
