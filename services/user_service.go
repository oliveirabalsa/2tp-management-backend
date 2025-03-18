package services

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/middleware"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func RegisterUser(user *models.User) error {
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	// Set default role
	user.Role = "user"

	err := repositories.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func AuthenticateUser(username, password string) (string, error) {
	user, err := repositories.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	return middleware.GenerateJWT(user.ID, user.Username, user.Role)
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
