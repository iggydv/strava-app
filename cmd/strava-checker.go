package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strava-app/internal/web/stravaauth"
	"strava-app/pkg/strava"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	s := &strava.Connector{
		Client:       client,
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURI:  "http://localhost:8080/callback",
	}

	// Initialize auth config
	config := stravaauth.StravaConfig{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURI:  "http://localhost:8080/callback",
	}

	// Token callback function to update the connector
	tokenHandler := func(tokens stravaauth.TokenResponse) error {
		s.SetTokens(tokens)
		// Here you might also want to persist tokens to a database
		return nil
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/strava", stravaauth.AuthHandler(config))
	mux.HandleFunc("/callback", stravaauth.CallbackHandler(config, tokenHandler))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Printf("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
