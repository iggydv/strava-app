package web

import "net/http"

type ServerInterface interface {
	Auth(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
	GetAthlete(w http.ResponseWriter, r *http.Request)
	GetAthleteStats(w http.ResponseWriter, r *http.Request)
	GetAthleteTotalDistance(w http.ResponseWriter, r *http.Request)
}
