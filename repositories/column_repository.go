package repositories

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
)

func CreateColumn(column *models.Column) error {
	return config.DB.Create(column).Error
}

func GetColumnByID(columnID uuid.UUID) (*models.Column, error) {
	var column models.Column
	err := config.DB.Where("id = ?", columnID).Preload("Tasks").First(&column).Error
	return &column, err
}

func GetColumnsByBoard(boardID uuid.UUID) ([]models.Column, error) {
	var columns []models.Column
	err := config.DB.Where("board_id = ?", boardID).Preload("Tasks").Preload("Tasks.Assigner").Find(&columns).Error
	return columns, err
}

func DeleteColumn(columnID uuid.UUID) error {
	return config.DB.Delete(&models.Column{}, columnID).Error
}
