package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/routes"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found. Using default values.")
	}

	// Connect to database
	config.ConnectDatabase()

	// Initialize router
	router := gin.Default()

	// Load routes
	routes.UserRoutes(router)
	routes.BoardRoutes(router)
	routes.ColumnRoutes(router)
	routes.TaskRoutes(router)

	// Start server
	router.Run(":8081")
}
