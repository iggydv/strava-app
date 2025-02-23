package web

import "net/http"

type ServerInterface interface {
	Auth(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
	GetAthlete(w http.ResponseWriter, r *http.Request)
}
