package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/controllers"
)

func ColumnRoutes(router *gin.Engine) {
	router.POST("/api/columns", controllers.CreateColumn)
	router.GET("/api/boards/:board_id/columns", controllers.GetColumnsByBoard)
	router.DELETE("/api/columns/:column_id", controllers.DeleteColumn)
}
