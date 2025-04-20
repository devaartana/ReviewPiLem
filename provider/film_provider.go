package provider

import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/repository"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideFilmDepedencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	filmRepository := repository.NewFilmRepository(db)
	genreRepository := repository.NewGenreRepository(db)
	filmImagesRepository := repository.NewFilmImagesRepository(db)
	reviewrepository := repository.NewReviewRepository(db) 

	filmService := service.NewFilmServices(filmRepository, genreRepository, filmImagesRepository, reviewrepository)

	do.Provide(injector, func (i *do.Injector) (controller.FilmController, error){
		return controller.NewFilmController(filmService), nil
	})
}