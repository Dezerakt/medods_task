package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"medods_task/cmd/middlewares"
	"medods_task/configs"
	"medods_task/internal"
)

var (
	config *configs.ConfigLoader
	server *gin.Engine

	routeHandler *internal.RouteHandler
)

func init() {
	server = gin.Default()
	config = configs.NewConfigLoader()
	routeHandler = internal.NewRouteHandler()

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, "POST")
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, "GET")

	server.Use(cors.New(corsConfig), middlewares.InitializerMiddleware())

	routeGroup := server.Group("/api")

	routeHandler.Handle(routeGroup)

	log.Fatal(server.Run(":" + config.EnvConfig.ServerPort))
}
