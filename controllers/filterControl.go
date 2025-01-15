package controllers

import (
	"net/http"
	"strconv"

	"github.com/MichaelSitanggang/bookstore/services"
	"github.com/gin-gonic/gin"
)

type FilterControl struct {
	servic services.FilterBookService
}

func NewFilterControl(servic services.FilterBookService) *FilterControl {
	return &FilterControl{servic: servic}
}

func (ctrl *FilterControl) CariBooks(c *gin.Context) {
	judul := c.Query("judul")
	tahun, _ := strconv.Atoi(c.Query("tahun"))
	if tahun <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ada"})
		return
	}
	filter, err := ctrl.servic.CariBooks(judul, tahun)
	if filter.ID == 0 && filter.Judul == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "code": 200, "meta": gin.H{"data": filter}})
}

func (ctrl *FilterControl) TampilaknByPenjualan(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "4"))
	book, err := ctrl.servic.TampilkanPenjualan(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "berhasil", "code": 200, "meta": gin.H{"data": book}})
}
