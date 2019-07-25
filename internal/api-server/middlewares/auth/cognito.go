package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/equest/exam/internal/auth"
	"github.com/equest/exam/pkg/errors"
)

func NewCognitoValidator(aws *auth.AWSCognitoAuthService) Validator {
	return &CognitoValidator{
		Auths: aws,
	}
}

// CognitoValidator implementation of
type CognitoValidator struct {
	Auths *auth.AWSCognitoAuthService
}

// Validate validates key against a store
func (m CognitoValidator) Validate(key string) (*Identity, error) {
	if key == "" {
		return nil, errors.NewAuthError("unauthorized")
	}
	jwtToken, err := m.Auths.ParseJWT(key)
	if err != nil {
		return nil, errors.NewAuthError("unauthorized")
	}
	if !jwtToken.Valid {
		return nil, errors.NewAuthError("unauthorized")
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	use := claims["token_use"].(string)
	if use != "access" {
		return nil, errors.NewAuthError("unauthorized")
	}
	username := claims["username"].(string)
	if username == "" {
		return nil, errors.NewAuthError("unauthorized")
	}
	return &Identity{
		Type: "aws-cognito",
		ID:   username,
		Name: username,
	}, nil
}
