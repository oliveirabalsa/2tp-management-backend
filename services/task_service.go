package services

import (
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func CreateTaskService(task *models.Task) error {
	return repositories.CreateTask(task)
}

func GetColumnTasks(columnID uint) ([]models.Task, error) {
	return repositories.GetTasksByColumn(columnID)
}

func DeleteTaskService(taskID uint) error {
	return repositories.DeleteTask(taskID)
}
