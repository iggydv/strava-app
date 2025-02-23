package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strava-app/internal/strava/web/models"
)

// this struct-based approach is a bit more verbose than the function-based approach, but it allows for more flexibility
// in the future. For example, the AuthHandler struct could be extended to include additional fields or methods.
// This is useful if the handler needs to maintain state or perform additional operations.
// Alternative approach:
// func AuthHandler(config StravaConfig) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) { ... }}

type AuthHandler struct {
	config        models.StravaConfig
	tokenCallback func(models.TokenResponse) error
}

func NewAuthHandler(config models.StravaConfig, tokenCallback func(models.TokenResponse) error) AuthHandler {
	return AuthHandler{
		config:        config,
		tokenCallback: tokenCallback,
	}
}

func (a *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authenticating")
	// Redirect to Strava OAuth page
	authURL := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=activity:read_all,profile:read_all",
		a.config.ClientID,
		a.config.RedirectURI)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (a *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Callback received")
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for tokens
	tokenURL := "https://www.strava.com/oauth/token"
	resp, err := http.PostForm(tokenURL, map[string][]string{
		"client_id":     {a.config.ClientID},
		"client_secret": {a.config.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	})
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse the token response
	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		http.Error(w, "Failed to parse token response", http.StatusInternalServerError)
		return
	}

	// Call the provided callback function with the token response
	if err := a.tokenCallback(tokenResp); err != nil {
		http.Error(w, "Failed to process token", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Authentication successful",
	})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
