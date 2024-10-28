package usersController

import (
	"encoding/json"
	"net/http"

	"github.com/italodavidb/goCrud/database"
	"github.com/italodavidb/goCrud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	database.DB.Create(&newUser)

	newUser.Password = ""

	json.NewEncoder(w).Encode(newUser)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	for i := range users {
		users[i].Password = ""
	}
	json.NewEncoder(w).Encode(users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
		return
	}
	user, err := FetchUser(username)
	if err != nil {
		HandleUserFetchError(w, err)
		return
	}
	user.Password = ""
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
	}
	if err := database.DB.Where("username = ?", username).Delete(&models.User{}).Error; err != nil {
		HandleUserFetchError(w, err)
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

	existingUser, err := FetchUser(username)
	if err != nil {
		HandleUserFetchError(w, err)
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

	existingUser.Password = "" // Não retornamos a senha
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func FetchUser(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func HandleUserFetchError(w http.ResponseWriter, err error) bool {
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return true
	}
	http.Error(w, "Error fetching user", http.StatusInternalServerError)
	return true
}
