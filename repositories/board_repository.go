package repositories

import (
	"fmt"

	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
)

func CreateBoard(board *models.Board) error {
	return config.DB.Create(board).Error
}

func GetBoards() ([]models.Board, error) {
	var boards []models.Board
	err := config.DB.Find(&boards).Error
	return boards, err
}

func BoardExists(boardID uint) bool {
	var board models.Board
	config.DB.First(&board, boardID)
	fmt.Println(board)
	return board.ID != 0
}
