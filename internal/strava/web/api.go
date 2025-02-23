package web

import (
	"strava-app/internal/strava/web/handlers"
)

var _ ServerInterface = &API{}

type API struct {
	handlers.AuthHandler
	handlers.AthleteHandler
}
