package stravaauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthHandler returns a handler function for Strava OAuth
func AuthHandler(config StravaConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling AuthHandler")
		// Redirect to Strava OAuth page
		authURL := fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=activity:read_all,profile:read_all",
			config.ClientID,
			config.RedirectURI)
		http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
	}
}

// CallbackHandler returns a handler function for the OAuth callback
func CallbackHandler(config StravaConfig, tokenCallback func(TokenResponse) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the authorization code from the URL parameters
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Missing authorization code", http.StatusBadRequest)
			return
		}

		// Exchange the authorization code for tokens
		tokenURL := "https://www.strava.com/oauth/token"
		resp, err := http.PostForm(tokenURL, map[string][]string{
			"client_id":     {config.ClientID},
			"client_secret": {config.ClientSecret},
			"code":          {code},
			"grant_type":    {"authorization_code"},
		})
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Parse the token response
		var tokenResp TokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
			http.Error(w, "Failed to parse token response", http.StatusInternalServerError)
			return
		}

		// Call the provided callback function with the token response
		if err := tokenCallback(tokenResp); err != nil {
			http.Error(w, "Failed to process token", http.StatusInternalServerError)
			return
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Authentication successful",
		})
	}
}
