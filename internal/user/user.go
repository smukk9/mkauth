package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Remove static test data
	users := []User{
		{Id: "1", Username: "alice"},
		{Id: "2", Username: "bob"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
