package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/bookstore/services"
	"github.com/gin-gonic/gin"
)

type RatingController struct {
	ratingService services.RatingServices
}

func NewRatingControl(ratingService services.RatingServices) *RatingController {
	return &RatingController{ratingService: ratingService}
}

func (ctrl *RatingController) AddsRating(c *gin.Context) {
	type request struct {
		BookID int    `json:"bookID"`
		Rating int    `json:"rating"`
		Ulasan string `json:"ulasan"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("accountID")
	if err := ctrl.ratingService.TambahReview(userID.(int), req.BookID, req.Rating, req.Ulasan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ulasan berhasil ditambah"})
}
