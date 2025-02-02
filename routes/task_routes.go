package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/controllers"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
)

func TaskRoutes(router *gin.Engine) {
	router.POST("/api/tasks", middleware.AuthMiddleware(), controllers.CreateTask)
	router.GET("/api/columns/:column_id/tasks", controllers.GetTasksByColumn)
	router.DELETE("/api/tasks/:task_id", controllers.DeleteTask)
}
