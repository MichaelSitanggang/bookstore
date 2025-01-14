package controllers

import (
	"net/http"
	"strconv"

	"github.com/MichaelSitanggang/bookstore/services"
	"github.com/gin-gonic/gin"
)

type BookControl struct {
	servic services.BookService
}

func NewBookControl(servic services.BookService) *BookControl {
	return &BookControl{servic: servic}
}

func (ctrl *BookControl) GetAllBooks(c *gin.Context) {
	books, err := ctrl.servic.GetAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "code": 200, "meta": gin.H{"data": books}})
}

func (ctrl *BookControl) GetByIdBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id param error"})
		return
	}
	books, err := ctrl.servic.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "code": 200, "meta": gin.H{"data": books}})
}

func (ctrl *BookControl) CreatedBooks(c *gin.Context) {
	type Inputs struct {
		Gambar string  `json:"gambar"`
		Judul  string  `json:"judul"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Harga  float64 `json:"harga"`
		Stok   int     `json:"stok"`
	}
	var inputan Inputs
	if err := c.ShouldBindJSON(&inputan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "inputan salah"})
		return
	}
	if err := ctrl.servic.CreateBook(inputan.Gambar, inputan.Judul, inputan.Author, inputan.Year, inputan.Harga, inputan.Stok); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "code": 200, "meta": gin.H{"message": "Data Berhasil Ditambah"}})
}
