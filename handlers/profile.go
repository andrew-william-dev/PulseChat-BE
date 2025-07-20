package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"chatapp/models"
	"chatapp/utils"
)

func getSafeString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var (
		username, email, bio, avatar, location, website, interests, theme sql.NullString
		user                                                   models.User
	)

	err := utils.DB.QueryRow(
		`SELECT id, username, email, bio, avatar_url, location, website, interests, theme_color
		 FROM users 
		 WHERE id = $1`,
		userID,
	).Scan(
		&user.ID,
		&username,
		&email,
		&bio,
		&avatar,
		&location,
		&website,
		&interests,
		&theme,
	)

	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Username = getSafeString(username)
	user.Email = getSafeString(email)
	user.Bio = getSafeString(bio)
	user.Avatar = getSafeString(avatar)
	user.Location = getSafeString(location)
	user.Website = getSafeString(website)
	user.Interests = getSafeString(interests)
	user.Theme = getSafeString(theme)


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": user,
	})
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var data struct {
		Bio       string `json:"bio"`
		Avatar    string `json:"avatar"`
		Location  string `json:"location"`
		Website   string `json:"website"`
		Interests   string `json:"interests"`
		Theme     string `json:"theme_color"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := utils.DB.Exec(
		`UPDATE users 
		 SET bio = $1, avatar_url = $2, location = $3, website = $4, interests = $5, theme_color = $6
		 WHERE id = $7`,
		data.Bio, data.Avatar, data.Location, data.Website, data.Interests, data.Theme, userID,
	)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

