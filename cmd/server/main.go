package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/kyma-go-service/internal/routes"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :8080")
	r.Run(":" + port)
}
