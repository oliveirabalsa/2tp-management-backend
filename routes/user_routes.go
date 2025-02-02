package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/controllers"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
)

func UserRoutes(router *gin.Engine) {
	public := router.Group("/api")
	{
		public.POST("/signup", controllers.Signup)
		public.POST("/login", controllers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controllers.Logout)
	}
}
