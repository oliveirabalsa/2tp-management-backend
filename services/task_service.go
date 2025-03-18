package services

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func CreateTaskService(task *models.Task) error {
	return repositories.CreateTask(task)
}

func GetColumnTasks(columnID uuid.UUID) ([]models.Task, error) {
	return repositories.GetTasksByColumn(columnID)
}

func DeleteTaskService(taskID uuid.UUID) error {
	return repositories.DeleteTask(taskID)
}

func UpdateTaskService(taskID uuid.UUID, updates map[string]interface{}) error {
	task, err := repositories.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	task.Title = updates["title"].(string)
	task.Description = updates["description"].(string)
	return repositories.UpdateTask(taskID, updates)
}

func GetTaskByID(taskID uuid.UUID) (*models.Task, error) {
	return repositories.GetTaskByID(taskID)
}
