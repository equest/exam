package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equest/exam/internal/api-server/handlers"
	"github.com/equest/exam/internal/api-server/middlewares/auth"
	"github.com/equest/exam/pkg/errors"
)

var handler = handlers.Handler(func(w http.ResponseWriter, r *http.Request) error {
	id := auth.FromContext(r.Context())
	if id == nil {
		return errors.NewAuthError("id nil")
	}
	return nil
})

func Test_StaticAPIKeyValidatorMiddleware_ShouldOK(t *testing.T) {
	middleware := auth.Middleware(auth.StaticAPIKeyValidator{})
	h := middleware(handler)
	server := httptest.NewServer(h)

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer my-secret-api-key")

	res, err := server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %v, got %v", http.StatusOK, res.StatusCode)
	}
}

func Test_StaticAPIKeyValidatorMiddleware_ShouldUnauthorized(t *testing.T) {
	middleware := auth.Middleware(auth.StaticAPIKeyValidator{})
	h := middleware(handler)
	server := httptest.NewServer(h)

	// req, err := http.NewRequest("GET", server.URL, nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// req.Header.Set("Authorization", "Bearer my-secret-api-key")

	res, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status %v, got %v", http.StatusUnauthorized, res.StatusCode)
	}
}
