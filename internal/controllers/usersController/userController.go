package usersController

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/internal/database"
	"github.com/italodavidb/goCrud/internal/models"
	"github.com/italodavidb/goCrud/internal/repository"
	"github.com/italodavidb/goCrud/internal/services"
	"github.com/italodavidb/goCrud/internal/utils/errorUtils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Erro ao processar a entrada", http.StatusBadRequest)
		return
	}

	createdUser, err := services.CreateUser(newUser)
	if err != nil {
		// Verifica o tipo de erro retornado
		switch err.Error() {
		case "usuário já existe":
			http.Error(w, "Usuário já existe", http.StatusConflict) // 409 Conflict
			return
		default:
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}
	}

	// Se a criação for bem-sucedida
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser) // Opcional: Retorna o novo usuário

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
	if err := database.DB.Where("username = ?", username).Delete(&models.User{}).Error; err != nil {
		errorUtils.HandleUserFetchError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	existingUser, err := repository.FetchUser(username)
	if err != nil {
		errorUtils.HandleUserFetchError(w, err)
		return
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if err := database.DB.Save(existingUser).Error; err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	existingUser.Password = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}
