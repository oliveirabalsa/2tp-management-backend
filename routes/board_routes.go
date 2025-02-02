package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/controllers"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
)

func BoardRoutes(router *gin.Engine) {
	protected := router.Group("/api/boards")
	{
		protected.POST("/", middleware.AuthMiddleware(), controllers.CreateBoard)
		protected.GET("/", controllers.GetBoards)
	}
}
