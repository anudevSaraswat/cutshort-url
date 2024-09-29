package routes

import (
	"github.com/anudevSaraswat/cutshort-url/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowMethods: []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Authorization", "Content-Type", "Options", "Accept", "Referer", "User-Agent", "Version"},
		AllowOrigins: []string{"*"},
	}))

	server.POST("/api/shorten", handlers.APIShortenURL)
	server.GET("/api/resolve/:short_url", handlers.APIResolveURL)

	return server

}
