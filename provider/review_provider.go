package provider

import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/repository"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/samber/do"
	"gorm.io/gorm"
)


func ProvideReviewDependencies(injector *do.Injector) {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Repository
	reviewRepository := repository.NewReviewRepository(db)
	reactionRepository := repository.NewReactionRepository(db)

	// Service
	reviewService := service.NewReviewServices(reviewRepository, reactionRepository, jwtService)

	// Controller
	do.Provide(injector, func (i *do.Injector) (controller.ReviewController, error){
		return controller.NewReviewController(reviewService), nil
	})
}