package cardsController

import (
	"encoding/json"
	"net/http"

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
	type CardPayload struct {
		SetCode string `json:"SetCode"`
		Number  string `json:"Number"`
	}

	var payload struct {
		Cards []CardPayload `json:"cards"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var cards []models.Card
	for _, card := range payload.Cards {
		if card.SetCode == "" || card.Number == "" {
			http.Error(w, "Each card must have a SetCode and Number", http.StatusBadRequest)
			return
		}
		cards = append(cards, models.Card{
			SetCode: card.SetCode,
			Number:  card.Number,
		})
	}

	err := services.DeleteCards(cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	setCode := r.URL.Query().Get("setCode")
	number := r.URL.Query().Get("number")

	var updatedCard models.Card
	if err := json.NewDecoder(r.Body).Decode(&updatedCard); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	card, err := services.UpdateCard(setCode, number, updatedCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}
