package handlers

import (
	"net/http"

	"github.com/equest/exam/internal/api-server/response"
	"github.com/equest/exam/internal/app"
	"github.com/equest/exam/pkg/alert"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Handler(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			return
		default:
			err := fn(w, r)
			if err != nil {
				app := app.FromContext(r.Context())
				var alert alert.Alert
				if app != nil {
					alert = app.Alert
				}
				response.WithError(w, alert, err)
			}
			return
		}
	}
}
