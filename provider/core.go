package provider

import (
	"github.com/devaartana/ReviewPiLem/config"
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func (i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func (i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	ProvideUserDependencies(injector)
	ProvideGenreDependencies(injector)
	ProvideFilmDepedencies(injector)
	ProvideReviewDependencies(injector)
	ProvideUserFilmDependencies(injector)
}