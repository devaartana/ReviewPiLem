package main

import (
	"log"
	"os"

	"github.com/devaartana/ReviewPiLem/command"
	"github.com/devaartana/ReviewPiLem/middleware"
	"github.com/devaartana/ReviewPiLem/provider"
	"github.com/devaartana/ReviewPiLem/routes"
	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func args(injector *do.Injector) bool {
	if len(os.Args) > 1 {
		flag := command.Commands(injector)
		return flag
	}

	return true
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	if utils.GetEnvBool("IS_LOGGER", true) {
		// routes.LoggerRoute(server)
	}

	port := utils.GetEnvString("PORT", "8000") 
	
	var serve string
	if utils.GetEnvString("APP_ENV", "localhost") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
	var (
		injector = do.New()
	)

	provider.RegisterDependencies(injector)

	if !args(injector) {
		return 
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	server.Use(middleware.CORSMiddleware())

	routes.RegisterRoutes(server, injector)

	run(server)
}