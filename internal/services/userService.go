package services

import (
	"errors"

	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/repository"
	"github.com/italodavidb/goCrud/internal/utils/hashUtils"
)

func CreateUser(newUser models.User) (models.User, error) {
	hashedPassword, err := hashUtils.HashPassword(newUser.Password)
	if err != nil {
		return models.User{}, err
	}
	newUser.Password = string(hashedPassword)

	err = repository.CreateUser(newUser)
	if err != nil {
		return models.User{}, err // Retorna erro se falhar na criação
	}
	newUser.Password = ""
	return newUser, nil

}
func FindAllUsers() ([]models.User, error) {
	users, err := repository.FindAllUsers()
	if err != nil {
		return []models.User{}, err
	}
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

func FindUser(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is necessary")
	}
	user, err := repository.FetchUser(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
