package controllers

import (
	"encoding/json"
	"muse-dashboard-api/models"
	"muse-dashboard-api/utilities"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){
	var credential models.CredentialsAuth

	decoder := json.NewDecoder(r.Body)

	if err:= decoder.Decode(&credential) ; err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userData, err := models.Login(credential); 
	
	if err != nil{
		http.Error(w, "Invalid Credential", http.StatusBadRequest)
		return
	}

	accessToken, err := utilities.GenerateAccessToken(userData.ID)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utilities.GenerateRefreshToken(userData.ID)
	if err != nil {
		http.Error(w, "Could not generate refresh token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"id":userData.ID,
		"email":userData.Email,
		"username":userData.Username,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	claims, err := utilities.ParseToken(request.RefreshToken, utilities.RefreshTokenSecret)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	// Generate a new access token
	accessToken, err := utilities.GenerateAccessToken(claims.UserID)
	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access_token": accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}