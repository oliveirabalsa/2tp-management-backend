package repositories

import (
	"github.com/oliveirabalsa/2tp-management-backend/config"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Role = "user"

	return config.DB.Create(user).Error
}

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
