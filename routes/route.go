package routes

import (
	"github.com/MichaelSitanggang/bookstore/controllers"
	"github.com/MichaelSitanggang/bookstore/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(authcontrol *controllers.AuthControl, bookcontrol *controllers.BookControl, filtercontrol *controllers.FilterControl) *gin.Engine {
	r := gin.Default()
	r.POST("/register", authcontrol.Registers)
	r.POST("/login", authcontrol.Login)
	r.POST("/otpverifikasi", authcontrol.VerifikasiOtps)
	route := r.Group("/")
	route.Use(middlewares.AuthJwt())
	route.GET("/books", bookcontrol.GetAllBooks)
	route.GET("/books/:id", bookcontrol.GetByIdBooks)
	route.POST("/books", bookcontrol.CreatedBooks)
	route.GET("/searchbooks", filtercontrol.CariBooks)
	route.GET("/book-terjual", filtercontrol.TampilaknByPenjualan)
	return r
}
