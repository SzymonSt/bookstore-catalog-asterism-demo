package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	setupTestServer()
}

func setupTestServer() {
	var (
		server = gin.Default()
	)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthz", func(ctx *gin.Context) {
		message := "devx api is up and running"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	go func() { log.Fatal(server.Run(":" + "8080")) }()
}
