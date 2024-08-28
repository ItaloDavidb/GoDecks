package cards_controllers

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/database"
	"github.com/italodavidb/goCrud/models"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var newCards []models.Card

	if err := json.NewDecoder(r.Body).Decode(&newCards); err != nil {
		var singleCard models.Card
		if err := json.NewDecoder(r.Body).Decode(&singleCard); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newCards = append(newCards, singleCard)
	}

	for _, card := range newCards {
		if err := database.DB.Create(&card).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(newCards)
}

func FindAllCards(w http.ResponseWriter, r *http.Request) {
	var c []models.Card
	database.DB.Find(&c)
	json.NewEncoder(w).Encode(c)
}

func FindCards(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	setCode := r.URL.Query().Get("set_code")
	number := r.URL.Query().Get("number")

	var cards []models.Card

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
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cards[0])
}

func DeleteCards(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Cards []string `json:"cards"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for _, setcodeNumber := range payload.Cards {
		setCode := setcodeNumber[:3]
		number := setcodeNumber[3:]

		database.DB.Where("set_code = ? AND number = ?", setCode, number).Delete(&models.Card{})
	}

	w.WriteHeader(http.StatusNoContent)
}
