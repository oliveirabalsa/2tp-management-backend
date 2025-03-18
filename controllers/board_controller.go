package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func CreateBoard(c *gin.Context) {
	type BoardInput struct {
		Title string `json:"title" binding:"required"`
	}

	var input BoardInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No user ID found"})
		return
	}

	// Get user to check role
	user, err := services.GetUserByID(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user details"})
		return
	}

	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create boards"})
		return
	}

	board := models.Board{
		Title:   input.Title,
		AdminID: userID.(uuid.UUID),
	}

	if err := services.CreateBoardService(&board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create board"})
		return
	}

	c.JSON(http.StatusCreated, board)
}

func GetBoards(c *gin.Context) {
	boards, err := services.GetAllBoards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch boards"})
		return
	}

	c.JSON(http.StatusOK, boards)
}
