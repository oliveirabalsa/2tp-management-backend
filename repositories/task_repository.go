package repositories

import (
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTasksByColumn(columnID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Preload("Assigner").Where("column_id = ?", columnID).Find(&tasks).Error
	return tasks, err
}

func DeleteTask(taskID uint) error {
	return config.DB.Delete(&models.Task{}, taskID).Error
}
