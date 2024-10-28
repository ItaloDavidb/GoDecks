package errorUtils

import (
	"net/http"

	"gorm.io/gorm"
)

func HandleUserFetchError(w http.ResponseWriter, err error) bool {
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return true
	}
	http.Error(w, "Error fetching user", http.StatusInternalServerError)
	return true
}
