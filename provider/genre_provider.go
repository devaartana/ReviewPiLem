package provider

import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/repository"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)


func ProvideGenreDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	genreRepository := repository.NewGenreRepository(db)

	// Service
	genreService := service.NewGenreService(genreRepository, jwtService)

	// Controller
	do.Provide(injector, func (i *do.Injector) (controller.GenreController, error){
		return controller.NewGenreController(genreService), nil
	})
}