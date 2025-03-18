package repositories

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
)

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func GetTasksByColumn(columnID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Where("column_id = ?", columnID).Find(&tasks).Error
	return tasks, err
}

func GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := config.DB.Where("id = ?", taskID).First(&task).Error
	return &task, err
}

func DeleteTask(taskID uuid.UUID) error {
	return config.DB.Delete(&models.Task{}, taskID).Error
}

func UpdateTask(taskID uuid.UUID, updates map[string]interface{}) error {
	return config.DB.Model(&models.Task{}).Where("id = ?", taskID).Updates(updates).Error
}
