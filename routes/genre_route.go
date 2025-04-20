package routes


import (

	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/middleware"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Genre(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	genreController := do.MustInvoke[controller.GenreController](injector)
	
	routes := route.Group("/api/genre")
	{
		routes.GET("", genreController.GetAllGenre)
		routes.POST("", middleware.AuthorizeAdmin(jwtService), genreController.Create)
		routes.PUT("/:id", middleware.AuthorizeAdmin(jwtService), genreController.Update)
	}
}