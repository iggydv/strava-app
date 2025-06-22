package strava

import (
	"encoding/json"
	"fmt"
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

func (c *Connector) GetAthlete() (*models.Athlete, int, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.strava.com/api/v3/athlete", nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to get athlete: %v", err)
	}

	defer resp.Body.Close()

	var athlete *models.Athlete
	if err := parse.JSON(resp.Body, &athlete); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return athlete, resp.StatusCode, nil
}

func (c *Connector) GetAthleteStats() (*models.AthleteStats, int, error) {
	// Get athlete information
	athlete, status, err := c.GetAthlete()
	if err != nil {
		return nil, status, err
	}
	if athlete.ID == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("athlete ID missing")
	}

	// Get athlete stats
	stats, status, err := c.getAthleteStats(athlete.ID)
	if err != nil {
		return nil, status, err
	}

	return stats, status, nil
}

func (c *Connector) getAthleteStats(athleteID int) (*models.AthleteStats, int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.strava.com/api/v3/athletes/%d/stats", athleteID), nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	var stats *models.AthleteStats
	if err := parse.JSON(resp.Body, &stats); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return stats, http.StatusOK, nil
}

func (c *Connector) SetTokens(tokens models.TokenResponse) {
	c.AccessToken = tokens.AccessToken
	c.RefreshedToken = tokens.RefreshToken
	c.ExpiresAt = tokens.ExpiresAt
}

func (c *Connector) RefreshIfNeeded() error {
	if !c.isTokenExpired() {
		return nil
	}
	return c.refreshToken()
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
