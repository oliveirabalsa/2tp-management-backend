package repositories

import (
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
)

func CreateColumn(column *models.Column) error {
	return config.DB.Create(column).Error
}

func GetColumnsByBoard(boardID uint) ([]models.Column, error) {
	var columns []models.Column
	err := config.DB.Where("board_id = ?", boardID).Preload("Tasks").Find(&columns).Error
	return columns, err
}

func DeleteColumn(columnID uint) error {
	return config.DB.Delete(&models.Column{}, columnID).Error
}
