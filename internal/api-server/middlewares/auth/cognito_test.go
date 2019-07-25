package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equest/exam/internal/test"

	"github.com/equest/exam/internal/api-server/middlewares/auth"
)

func Test_CognitoValidatorMiddleware_ShouldOK(t *testing.T) {
	middleware := auth.Middleware(auth.CognitoValidator{
		Auths: test.GetServices().Auths,
	})
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
