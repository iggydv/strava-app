package middleware

import (
	"net/http"
	"strava-app/internal/strava"
)

func TokenRefreshMiddleware(connector *strava.Connector) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Exclude the auth endpoint
			if r.URL.Path == "/auth" || r.URL.Path == "/callback" {
				next.ServeHTTP(w, r)
				return
			}

			if err := connector.RefreshIfNeeded(); err != nil {
				http.Error(w, "Failed to refresh token", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
