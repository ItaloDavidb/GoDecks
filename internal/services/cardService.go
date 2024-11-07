package services

import (
	"fmt"
	"strings"

	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/repository"
)

func CreateCard(cards []models.Card) ([]models.Card, error) {
	var failedCards []models.Card
	var errorMessages []string
	var successfullySavedCards []models.Card

	for _, card := range cards {
		err := repository.CreateCard(card)
		if err != nil {
			failedCards = append(failedCards, card)
			errorMessages = append(errorMessages, fmt.Sprintf("Erro ao salvar card %v: %v", card, err))
		} else {
			successfullySavedCards = append(successfullySavedCards, card)
		}
	}

	if len(failedCards) > 0 {
		return successfullySavedCards, fmt.Errorf("falhas ao salvar os cards: %v", strings.Join(errorMessages, ", "))
	}

	return successfullySavedCards, nil
}

func FindAllCards() ([]models.Card, error) {
	cards, err := repository.FindAllCards()
	if err != nil {
		return []models.Card{}, err
	}
	return cards, nil
}

func FindCards(name, setCode, number string) (*models.Card, error) {
	if name == "" && setCode == "" && number == "" {
		return nil, fmt.Errorf("nenhum parâmetro de filtro foi fornecido")
	}
	cards, err := repository.FindCards(name, setCode, number)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

// func DeleteCards()
