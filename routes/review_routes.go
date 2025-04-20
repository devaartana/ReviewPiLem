package routes

import (
	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/middleware"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Review(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	reviewController := do.MustInvoke[controller.ReviewController](injector)

	routes := route.Group("/api/review")
	{
		routes.GET("/:id", reviewController.GetAllReviewFilm)
		routes.POST("", middleware.Authenticate(jwtService), reviewController.Create)
		routes.PUT("", middleware.Authenticate(jwtService), reviewController.Update)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), reviewController.Delete)
		routes.GET("/:id/like", middleware.Authenticate(jwtService), reviewController.Like)
		routes.GET("/:id/dislike", middleware.Authenticate(jwtService), reviewController.Dislike)
		routes.DELETE("/:id/reaction", middleware.Authenticate(jwtService), reviewController.Delete)
	}
}
