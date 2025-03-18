package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func Signup(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("JSON binding error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	if err := services.RegisterUser(&user); err != nil {
		fmt.Printf("Error registering user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials format"})
		return
	}

	// Trim any whitespace
	credentials.Username = strings.TrimSpace(credentials.Username)
	credentials.Password = strings.TrimSpace(credentials.Password)

	// Validate input
	if credentials.Username == "" || credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

		credentials.Username, len(credentials.Password))

	token, err := services.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
