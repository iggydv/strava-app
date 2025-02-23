package handlers

import (
	"encoding/json"
	"net/http"
	"strava-app/internal/strava"
	"strava-app/internal/strava/athletes"
)

type AthleteHandler struct {
	client *strava.Connector
}

func NewAthleteHandler(client *strava.Connector) AthleteHandler {
	return AthleteHandler{
		client: client,
	}
}

func (a *AthleteHandler) GetAthlete(w http.ResponseWriter, _ *http.Request) {
	athlete, err, statusCode := a.client.GetAthlete()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&athlete)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (a *AthleteHandler) GetAthleteStats(w http.ResponseWriter, r *http.Request) {
	stats, err, code := a.client.GetAthleteStats()
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&stats)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (a *AthleteHandler) GetAthleteTotalDistance(w http.ResponseWriter, r *http.Request) {
	stats, err, code := a.client.GetAthleteStats()
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]float64{
		"total_distance": athletes.CalculateTotalDistance(stats),
	})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
