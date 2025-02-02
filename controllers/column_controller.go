package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/services"
)

func CreateColumn(c *gin.Context) {
	var column models.Column
	if err := c.ShouldBindJSON(&column); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateColumnService(&column); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create column"})
		return
	}

	c.JSON(http.StatusCreated, column)
}

func GetColumnsByBoard(c *gin.Context) {
	boardID, err := strconv.ParseUint(c.Param("board_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	columns, err := services.GetBoardColumns(uint(boardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch columns"})
		return
	}

	c.JSON(http.StatusOK, columns)
}

func DeleteColumn(c *gin.Context) {
	columnID, err := strconv.ParseUint(c.Param("column_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid column ID"})
		return
	}

	if err := services.DeleteColumnService(uint(columnID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete column"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Column deleted successfully"})
}
