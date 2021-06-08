package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/prabhatsharma/mango/pkg/routes"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "6080"
	}

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	// r := gin.Default()
	// Creates a router without any middleware by default
	r := gin.New()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	routes.SetRoutes(r) // Set up all API routes.

	fmt.Println("Environment is: " + os.Getenv("APP_ENV"))

	r.Run(":" + PORT)
	fmt.Println("Listeninig on port: ", PORT)
}
