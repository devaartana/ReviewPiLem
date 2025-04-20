package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	User(server, injector)
	Genre(server, injector)
	Film(server, injector)
	Review(server, injector)
	UserFilm(server, injector)
}
