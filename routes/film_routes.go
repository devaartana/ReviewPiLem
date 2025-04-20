package routes

import (
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Film(route *gin.Engine, injector *do.Injector) {
	filmController := do.MustInvoke[controller.FilmController](injector)

	routes := route.Group("/api/film") 
	{
		routes.GET("/:id", filmController.GetFilmDetail)
		routes.GET("", filmController.GetFilm)
		routes.GET("/image/:id", filmController.GetImage)
		
	}
}