package main

import (
	"github.com/MichaelSitanggang/bookstore/config"
	"github.com/MichaelSitanggang/bookstore/controllers"
	"github.com/MichaelSitanggang/bookstore/repositories"
	"github.com/MichaelSitanggang/bookstore/routes"
	"github.com/MichaelSitanggang/bookstore/services"
)

func main() {
	db := config.CreateDatabase()

	//repo
	repoAuth := repositories.NewAuthRepo(db)

	//services
	usecaseAuth := services.NewAuthService(repoAuth)

	//controllers
	controlAuth := controllers.NewAuthControl(usecaseAuth)

	r := routes.Router(controlAuth)
	r.Run(":8080")
}
