package controllers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"gouth/models"
	"net/http"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// For now we just return the user without saving
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered",
		"user":    user,
	})
}
