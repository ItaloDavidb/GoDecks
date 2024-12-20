package repository

import (
	"github.com/italodavidb/goCrud/internal/database"
	"github.com/italodavidb/goCrud/internal/models"
)

func CreateCard(card models.Card) error {
	result := database.DB.Create(&card)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func FindAllCards() ([]models.Card, error) {
	var cards []models.Card
	result := database.DB.Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}
	return cards, nil
}
func FindCards(name, setCode, number string) (*models.Card, error) {
	var cards models.Card
	query := database.DB.Model(&models.Card{})

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if setCode != "" {
		query = query.Where("set_code = ?", setCode)
	}
	if number != "" {
		query = query.Where("number = ?", number)
	}

	result := query.Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cards, nil
}
func FindCardBySetCodeAndNumber(setCode string, number string) (models.Card, error) {
	var card models.Card

	query := database.DB.Model(&models.Card{})
	if setCode != "" {
		query = query.Where("set_code = ?", setCode)
	}
	if number != "" {
		query = query.Where("number = ?", number)
	}

	result := query.First(&card)
	return card, result.Error
}

func SaveCard(card *models.Card) error {
	result := database.DB.Save(card)
	return result.Error
}
func DeleteCardBySetCodeAndNumber(setCode string, number string) error {
	result := database.DB.Where("set_code = ? AND number = ?", setCode, number).Delete(&models.Card{})
	return result.Error
}
