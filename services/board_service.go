package services

import (
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func CreateBoardService(board *models.Board) error {
	return repositories.CreateBoard(board)
}

func GetAllBoards() ([]models.Board, error) {
	return repositories.GetBoards()
}
