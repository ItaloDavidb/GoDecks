package repository

import (
	"github.com/italodavidb/goCrud/internal/database"
	"github.com/italodavidb/goCrud/internal/models"
)

func FetchUser(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(newuser models.User) error {
	result := database.DB.Create(&newuser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
