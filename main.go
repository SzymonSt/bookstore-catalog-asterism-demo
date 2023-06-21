package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		server = gin.Default()
	)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/bookstore-catalog/api")
	router.GET("/healthz", func(ctx *gin.Context) {
		message := "bookstore-catalogv2 api is up and running"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + "8080"))
}
