package repositories

import (
	"errors"
	"fmt"

	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	var existingUser models.User
	result := config.DB.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == nil {
		return errors.New("username already exists")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user.Password = string(hashedPassword)

	user.Role = "user"

	if err := config.DB.Create(user).Error; err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	return nil
}

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
