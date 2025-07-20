package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"chatapp/models"
	"chatapp/utils"
)

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	var userID int
	err := utils.DB.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		req.Username, req.Email, string(hash),
	).Scan(&userID)

	if err != nil {
		http.Error(w, "Registration failed", http.StatusBadRequest)
		return
	}

	token, _ := utils.GenerateToken(userID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	var user models.User
	err := utils.DB.QueryRow("SELECT id, password FROM users WHERE username = $1", req.Username).
		Scan(&user.ID, &user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("Password mismatch")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
