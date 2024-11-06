package services

import (
	"errors"
	"fmt"

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
		return models.User{}, err
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

func DeleteUser(username string) error {
	if err := repository.DeleteUser(username); err != nil {
		return err
	}
	return nil
}

func UpdateUser(username string, user models.User) (models.User, error) {
	existingUser, err := FindUser(username)
	if err != nil {
		return models.User{}, fmt.Errorf("usuário não encontrado: %w", err)
	}
	updated := false
	if user.Username != "" && user.Username != existingUser.Username {
		existingUser.Username = user.Username
		updated = true
	}
	if user.Email != "" && user.Email != existingUser.Email {
		existingUser.Email = user.Email
		updated = true
	}

	if !updated {
		return *existingUser, nil
	}

	updatedUser, err := repository.UpdateUser(existingUser)
	if err != nil {
		return models.User{}, fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return *updatedUser, nil
}
