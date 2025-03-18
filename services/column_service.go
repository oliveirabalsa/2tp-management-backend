package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func CreateColumnService(column *models.Column) error {
	boardExists := repositories.BoardExists(column.BoardID)

	if !boardExists {
		return errors.New("Board does not exist")
	}
	return repositories.CreateColumn(column)
}

func GetBoardColumns(boardID uuid.UUID) ([]models.Column, error) {
	return repositories.GetColumnsByBoard(boardID)
}

func DeleteColumnService(columnID uuid.UUID) error {
	return repositories.DeleteColumn(columnID)
}

func GetColumnByID(columnID uuid.UUID) (*models.Column, error) {
	return repositories.GetColumnByID(columnID)
}
