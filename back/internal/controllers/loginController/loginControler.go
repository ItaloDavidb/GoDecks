package loginController

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tokenString, err := services.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
