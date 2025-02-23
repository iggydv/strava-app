package strava

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strava-app/internal/parse"
	"strava-app/internal/strava/web/models"
	"time"
)

type Connector struct {
	Client         *http.Client
	AccessToken    string
	RefreshedToken string
	ExpiresAt      int64
	ClientID       string
	ClientSecret   string
	RedirectURI    string
}

func (c *Connector) GetAthlete() (*models.Athlete, error, int) {
	if err := c.refreshIfNeeded(); err != nil {
		return nil, fmt.Errorf("no refresh token available"), http.StatusUnauthorized
	}

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := c.Client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to get athlete: %s", resp.Status), resp.StatusCode
	}

	var athlete *models.Athlete
	if err := parse.JSON(resp.Body, &athlete); err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return athlete, nil, resp.StatusCode
}

func (c *Connector) GetAthleteStats() (*models.AthleteStats, error, int) {
	// Get athlete information
	athlete, err, status := c.GetAthlete()
	if err != nil {
		return nil, err, status
	}
	if athlete.ID == 0 {
		return nil, fmt.Errorf("athlete ID missing"), http.StatusNotFound
	}

	// Get athlete stats
	stats, err, status := c.getAthleteStats(athlete.ID)
	if err != nil {
		return nil, err, status
	}

	return stats, nil, http.StatusOK
}

func (c *Connector) getAthleteStats(athleteID int) (*models.AthleteStats, error, int) {
	if err := c.refreshIfNeeded(); err != nil {
		return nil, fmt.Errorf("no refresh token available"), http.StatusUnauthorized
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", athleteID), nil)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err, resp.StatusCode
	}
	defer resp.Body.Close()

	var stats *models.AthleteStats
	if err := parse.JSON(resp.Body, &stats); err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return stats, nil, http.StatusOK
}

func (c *Connector) SetTokens(tokens models.TokenResponse) {
	c.AccessToken = tokens.AccessToken
	c.RefreshedToken = tokens.RefreshToken
	c.ExpiresAt = tokens.ExpiresAt
}

func (c *Connector) refreshToken() error {
	if c.RefreshedToken == "" {
		return fmt.Errorf("no refresh token available")
	}

	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)
	data.Set("refresh_token", c.RefreshedToken)
	data.Set("grant_type", "refresh_token")

	resp, err := c.Client.PostForm("https://www.strava.com/oauth/token", data)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("refresh token request failed with status: %d", resp.StatusCode)
	}

	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return fmt.Errorf("failed to decode refresh response: %w", err)
	}

	c.AccessToken = tokenResp.AccessToken
	c.RefreshedToken = tokenResp.RefreshToken
	c.ExpiresAt = tokenResp.ExpiresAt

	return nil
}

func (c *Connector) isTokenExpired() bool {
	return time.Now().Unix() > c.ExpiresAt
}

func (c *Connector) refreshIfNeeded() error {
	if !c.isTokenExpired() {
		return nil
	}
	// Refresh token logic here
	return c.refreshToken()
}
