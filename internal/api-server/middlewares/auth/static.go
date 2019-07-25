package auth

import (
	"github.com/equest/exam/pkg/errors"
)

// StaticAPIKeyValidator implementation of
type StaticAPIKeyValidator struct {
}

// Validate validates key against a store
func (m StaticAPIKeyValidator) Validate(key string) (*Identity, error) {
	if key == "" {
		return nil, errors.NewAuthError("unauthorized")
	}
	return &Identity{
		Type: "static",
		ID:   key,
		Name: "shared-secret",
	}, nil
}
