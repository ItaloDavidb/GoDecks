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
