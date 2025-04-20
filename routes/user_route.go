package routes

import (

	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/controller"
	"github.com/devaartana/ReviewPiLem/middleware"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func User(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	userController := do.MustInvoke[controller.UserController](injector)
	
	routes := route.Group("/api/user")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("", userController.GetAllUser)
		routes.GET("/:username", userController.GetUser)
		// routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
		// routes.PATCH("", middleware.Authenticate(jwtService), userController.Update)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
	}
}
