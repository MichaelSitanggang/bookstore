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
	repoBook := repositories.NewBookRepo(db)

	//services
	usecaseAuth := services.NewAuthService(repoAuth)
	usecaseBook := services.NewBookService(repoBook)

	//controllers
	controlAuthStruck := controllers.NewAuthControl(usecaseAuth)
	controlBookStruck := controllers.NewBookControl(usecaseBook)

	r := routes.Router(controlAuthStruck, controlBookStruck)
	r.Run(":8080")
}
