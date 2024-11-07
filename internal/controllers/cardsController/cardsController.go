package cardsController

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/internal/database"
	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/services"
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
	cards, err := services.CreateCard(newCards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cards)
}

func FindAllCards(w http.ResponseWriter, r *http.Request) {
	var cards []models.Card
	cards, err := services.FindAllCards()
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(cards)
}

func FindCards(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	setCode := r.URL.Query().Get("setCode")
	number := r.URL.Query().Get("number")

	cards, err := services.FindCards(name, setCode, number)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&cards)
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

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	setCode := r.URL.Query().Get("setCode")
	number := r.URL.Query().Get("number")

	var card models.Card

	query := database.DB.Model(&models.Card{})

	if setCode != "" {
		query = query.Where("set_code = ?", setCode)
	}
	if number != "" {
		query = query.Where("number = ?", number)
	}

	result := query.Find(&card)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	database.DB.Save(&card)

	json.NewEncoder(w).Encode(card)
}
