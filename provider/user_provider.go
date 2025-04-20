package provider

import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/repository"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)


func ProvideUserDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	userRepository := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepository, jwtService)

	// Controller
	do.Provide(injector, func (i *do.Injector) (controller.UserController, error){
		return controller.NewUserController(userService), nil
	})
}