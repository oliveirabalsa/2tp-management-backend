package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func CreateBoard(c *gin.Context) {
	var board models.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	fmt.Printf("Received Board Data: %+v\n", board)

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No user ID found"})
		return
	}

	board.AdminID = userID.(uint)

	fmt.Println("Assigned Admin ID:", board.AdminID)

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
