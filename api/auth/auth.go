package auth

import (
	"context"
	"fmt"

	"github.com/go-faster/errors"
	"google.golang.org/api/idtoken"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GoogleID      string `json:"sub"`
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
	payload, err := idtoken.Validate(context.Background(), token, v.clientID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate token")
	}

	// Verify the audience is allowed
	audience := payload.Claims["aud"].(string)
	isAllowed := false
	for _, id := range v.allowedExtensionIDs {
		if audience == id {
			isAllowed = true
			break
		} else {
			return nil, fmt.Errorf("not allowed audience with id: %s tried to get access to the api", id)
		}
	}

	if !isAllowed {
		return nil, fmt.Errorf("token audience %s is not allowed", audience)
	}

	claims := &GoogleClaims{
		Email:         payload.Claims["email"].(string),
		EmailVerified: payload.Claims["email_verified"].(bool),
		Name:          payload.Claims["name"].(string),
		Picture:       payload.Claims["picture"].(string),
		GoogleID:      payload.Claims["sub"].(string),
	}

	return claims, nil
}
