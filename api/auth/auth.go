package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GoogleID      string `json:"sub"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

type TokenValidator struct {
	clientID            string
	allowedExtensionIDs []string
}

func NewTokenValidator(clientID string, allowedExtensionIDs []string) *TokenValidator {
	return &TokenValidator{
		clientID:            clientID,
		allowedExtensionIDs: allowedExtensionIDs,
	}
}

func (v *TokenValidator) ValidateGoogleToken(token string) (*GoogleClaims, error) {
	// Call the same endpoint as the frontend
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid token: status %d", resp.StatusCode)
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	claims := &GoogleClaims{
		Email:         userInfo.Email,
		EmailVerified: userInfo.VerifiedEmail,
		Name:          userInfo.Name,
		Picture:       userInfo.Picture,
		GoogleID:      userInfo.ID,
	}

	return claims, nil
}
