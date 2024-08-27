package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/database"
	"github.com/italodavidb/goCrud/models"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var newCard models.Card
	json.NewDecoder(r.Body).Decode(&newCard)
	database.DB.Create(&newCard)
	json.NewEncoder(w).Encode(newCard)
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
