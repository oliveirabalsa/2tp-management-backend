package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/oliveirabalsa/2tp-management-backend/middleware"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"github.com/oliveirabalsa/2tp-management-backend/repositories"
)

func RegisterUser(user *models.User) error {
	user.Role = "user"
	return repositories.CreateUser(user)
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
