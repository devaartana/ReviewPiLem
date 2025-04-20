package routes

import (

	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/middleware"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func UserFilm(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	
	userFilmController := do.MustInvoke[controller.UserFilmController](injector)
	
	routes := route.Group("/api/user-film/")
	{
		routes.POST("", middleware.Authenticate(jwtService),userFilmController.Create)
		routes.PUT("", middleware.Authenticate(jwtService),userFilmController.Update)
		routes.DELETE("/:id",middleware.Authenticate(jwtService), userFilmController.Delete)
		routes.GET("/:id", userFilmController.GetUserList)
	}
}
