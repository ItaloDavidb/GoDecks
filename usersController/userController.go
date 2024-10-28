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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
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
	var user models.User
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
	}
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching user", http.StatusInternalServerError)
		}
		return
	}
	user.Password = ""
	json.NewEncoder(w).Encode(user)

}
