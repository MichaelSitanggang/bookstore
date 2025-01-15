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
	repoFilter := repositories.NewFilterBook(db)

	//services
	usecaseAuth := services.NewAuthService(repoAuth)
	usecaseBook := services.NewBookService(repoBook)
	usecaseFilter := services.NewFilterService(repoFilter)

	//controllers
	controlAuthStruck := controllers.NewAuthControl(usecaseAuth)
	controlBookStruck := controllers.NewBookControl(usecaseBook)
	controlFilterStruck := controllers.NewFilterControl(usecaseFilter)

	r := routes.Router(controlAuthStruck, controlBookStruck, controlFilterStruck)
	r.Run(":8080")
}
