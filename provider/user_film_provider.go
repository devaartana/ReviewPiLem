package provider
import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/repository"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)


func ProvideUserFilmDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	userFilmRepository := repository.NewUserFilmRepository(db)

	// Service
	userFilmService := service.NewUserFilmServices(userFilmRepository, jwtService)

	// Controller
	do.Provide(injector, func (i *do.Injector) (controller.UserFilmController, error){
		return controller.NewUserFilmController(userFilmService), nil
	})
}