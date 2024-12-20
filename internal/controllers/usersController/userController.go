package usersController

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Erro ao processar a entrada", http.StatusBadRequest)
		return
	}

	createdUser, err := services.CreateUser(newUser)
	if err != nil {
		switch err.Error() {
		case "usuário já existe":
			http.Error(w, "Usuário já existe", http.StatusConflict)
			return
		default:
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	users, err := services.FindAllUsers()
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	user, err := services.FindUser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
	}
	err := services.DeleteUser(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedUser, err := services.UpdateUser(username, user)
	updatedUser.Password = ""
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}
