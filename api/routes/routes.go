package routes

import (
	"github.com/anudevSaraswat/cutshort-url/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	server := gin.Default()

	server.GET("/api/shorten", handlers.APIShortenURL)
	server.GET("/api/resolve/:short_url", handlers.APIResolveURL)

	return server

}
