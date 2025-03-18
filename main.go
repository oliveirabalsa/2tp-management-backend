package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found. Using default values.")
	}

	config.ConnectDatabase()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	routes.UserRoutes(router)
	routes.BoardRoutes(router)
	routes.ColumnRoutes(router)
	routes.TaskRoutes(router)

	fmt.Println("Server running on port 8081")
	router.Run(":8081")
}
