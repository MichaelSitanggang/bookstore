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
	repoRating := repositories.NewRatingRepo(db)

	//services
	usecaseAuth := services.NewAuthService(repoAuth)
	usecaseBook := services.NewBookService(repoBook)
	usecaseFilter := services.NewFilterService(repoFilter)
	usecaseRating := services.NewServicesRating(repoRating)

	//controllers
	controlAuthStruck := controllers.NewAuthControl(usecaseAuth)
	controlBookStruck := controllers.NewBookControl(usecaseBook)
	controlFilterStruck := controllers.NewFilterControl(usecaseFilter)
	controlRatingStruck := controllers.NewRatingControl(usecaseRating)

	r := routes.Router(controlAuthStruck, controlBookStruck, controlFilterStruck, controlRatingStruck)
	r.Run(":8080")
}
 