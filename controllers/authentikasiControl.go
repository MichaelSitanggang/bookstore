package controllers

import (
	"net/http"

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
		NamaLengkap string `json:"namalengkap"`
		Email       string `json:"email"`
		Umur        int    `json:"umur"`
		Password    string `json:"password"`
	}
	var inputan input
	if err := c.ShouldBindJSON(inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	_, err := ctrl.usecase.Register(inputan.NamaLengkap, inputan.Email, inputan.Password, inputan.Umur)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal register"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil registrasi"})
}

func (ctrl *AuthControl) VerifikasiOtps(c *gin.Context) {
	var inputan struct {
		Otp string `json:"otp"`
	}
	if err := c.ShouldBindJSON(inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	if err := ctrl.usecase.VerifikasiOtp(inputan.Otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal verifikasi otp"})
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
	if err := c.ShouldBindJSON(inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Inputan salah"})
		return
	}
	_, err := ctrl.usecase.Login(inputan.Email, inputan.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login berhasil"})

}
