package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/bookstore/entities"
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
	var rating entities.Rating

	userID, _ := c.Get("accountID")
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rating.UserID = userID.(int)
	err := ctrl.ratingService.TambangRating(rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.ratingService.UbahBookRating(rating.BookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rating added and book rating updated"})
}
