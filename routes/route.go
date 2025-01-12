package routes

import (
	"github.com/MichaelSitanggang/bookstore/controllers"
	"github.com/MichaelSitanggang/bookstore/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(authcontrol *controllers.AuthControl) *gin.Engine {
	r := gin.Default()
	r.POST("/register", authcontrol.Registers)
	r.POST("/login", authcontrol.Login)
	r.POST("/otpverifikasi", authcontrol.VerifikasiOtps)
	route := r.Group("/")
	route.Use(middlewares.AuthJwt())
	return r
}
