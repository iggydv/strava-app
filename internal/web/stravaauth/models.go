package stravaauth

type StravaConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	TokenType    string `json:"token_type"`
}
