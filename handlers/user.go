package handlers

import (
	"chatapp/models"
	"chatapp/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	if userIDVal == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	rows, err := utils.DB.Query("SELECT id, username, avatar_url FROM users")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
	var (
		id       int
		username string
		avatar   sql.NullString
	)
	if err := rows.Scan(&id, &username, &avatar); err != nil {
		http.Error(w, "Scan error", http.StatusInternalServerError)
		return
	}
	users = append(users, models.User{
		ID:       id,
		Username: username,
		Avatar:   getSafeString(avatar),
	})
}

	json.NewEncoder(w).Encode(users)
}
