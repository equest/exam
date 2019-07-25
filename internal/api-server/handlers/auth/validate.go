package auth

import (
	"net/http"

	"github.com/equest/exam/internal/api-server/middlewares/auth"

	"github.com/equest/exam/internal/api-server/handlers"
	"github.com/equest/exam/internal/api-server/response"
)

// Validate validate
func Validate() http.HandlerFunc {
	return handlers.Handler(func(w http.ResponseWriter, r *http.Request) error {
		id := auth.FromContext(r.Context())
		response.JSON(w, http.StatusOK, id)
		return nil
	})
}
