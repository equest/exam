package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/equest/exam/internal/api-server/handlers"
)

// Identity represents the api caller
type Identity struct {
	Type string
	ID   interface{}
	Name string
}

// Validator is a wrapper interface for validating request based on api key
type Validator interface {
	Validate(credentials string) (*Identity, error)
}

func Middleware(v Validator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) error {
			key := getAPIKey(r)
			id, err := v.Validate(key)
			if err != nil {
				return err
			}
			ctx := ToContext(r.Context(), id)
			next.ServeHTTP(w, r.WithContext(ctx))
			return nil
		}
		return handlers.Handler(fn)
	}
}

func getAPIKey(r *http.Request) string {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, " ")

	if len(splitToken) < 2 {
		return ""
	}

	token = strings.Trim(splitToken[1], " ")
	return token
}

type key string

const authKey = key("api-key")

// FromContext get app from context
func FromContext(ctx context.Context) *Identity {
	id, ok := ctx.Value(authKey).(*Identity)
	if !ok {
		return nil
	}
	return id
}

// ToContext put app to context
func ToContext(ctx context.Context, id *Identity) context.Context {
	return context.WithValue(ctx, authKey, id)
}
