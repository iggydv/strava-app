package main

import (
	"log"
	"net/http"
	"os"
	"strava-app/internal/strava"
	"strava-app/internal/strava/web/handlers"
	"strava-app/internal/strava/web/models"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := Config{
		BaseURL:      os.Getenv("BASE_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	if config.ClientID == "" || config.ClientSecret == "" {
		log.Fatal("Missing CLIENT_ID or CLIENT_SECRET")
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}
	s := &strava.Connector{
		Client:       httpClient,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURI:  "http://localhost:8080/callback",
	}

	// Initialize auth config
	stravaConfig := models.StravaConfig{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURI:  "http://localhost:8080/callback",
	}

	// Token callback function to update the connector
	tokenCallbackHandler := func(tokens models.TokenResponse) error {
		s.SetTokens(tokens)
		// Could save tokens to a database here
		return nil
	}

	stravaAuthHandler := handlers.NewAuthHandler(stravaConfig, tokenCallbackHandler)
	stravaAthleteHandler := handlers.NewAthleteHandler(s)
	r := chi.NewRouter()

	// TODO: Use generated OpenAPI server interface
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Get("/auth", stravaAuthHandler.Auth)
		r.Get("/callback", stravaAuthHandler.Callback)
		r.Get("/athlete", stravaAthleteHandler.GetAthlete)
		r.Get("/athlete/stats", stravaAthleteHandler.GetAthleteStats)
		r.Get("/athlete/distance/total", stravaAthleteHandler.GetAthleteTotalDistance)
	})

	if config.BaseURL == "" {
		log.Fatal("Missing BASE_URL")
	}

	if err := http.ListenAndServe(config.BaseURL, r); err != nil {
		log.Fatal(err)
	}
}
