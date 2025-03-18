package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func CreateColumn(c *gin.Context) {
	var input struct {
		Title   string    `json:"title" binding:"required"`
		BoardID uuid.UUID `json:"board_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	column := models.Column{
		Title:   input.Title,
		BoardID: input.BoardID,
	}

	if err := services.CreateColumnService(&column); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create column"})
		return
	}

	c.JSON(http.StatusCreated, column)
}

func GetColumnsByBoard(c *gin.Context) {
	boardID, err := uuid.Parse(c.Param("board_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	columns, err := services.GetBoardColumns(boardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch columns"})
		return
	}

	c.JSON(http.StatusOK, columns)
}

func DeleteColumn(c *gin.Context) {
	columnUUID, err := uuid.Parse(c.Param("column_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column ID format"})
		return
	}

	if err := services.DeleteColumnService(columnUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete column"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Column deleted successfully"})
}
