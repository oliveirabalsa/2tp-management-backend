package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Role = "user"

	if err := services.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
	tokenString, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		return
	}

	expTime, exists := c.Get("exp")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not retrieve token expiration"})
		return
	}

	expirationTime, ok := expTime.(time.Time)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid expiration time format"})
		return
	}

	middleware.BlacklistToken(tokenString.(string), expirationTime)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
