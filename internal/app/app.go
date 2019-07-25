package app

import (
	"context"
	"net/http"

	"github.com/equest/exam/internal/auth"

	"github.com/equest/exam/internal/author"
	"github.com/equest/exam/internal/core"
	"github.com/equest/exam/internal/graph"
	"github.com/equest/exam/internal/organization"
	"github.com/equest/exam/internal/question"
	"github.com/equest/exam/internal/subject"
	"github.com/equest/exam/pkg/alert"
)

type App struct {
	Services *Services
	Clients  *Clients
	Alert    alert.Alert
}

type Services struct {
	Auths         *auth.AWSCognitoAuthService
	Graph         *graph.Service
	Core          core.Service
	Questions     question.Service
	Subjects      subject.Service
	Authors       author.Service
	Organizations organization.Service
}

type Clients struct {
}

type key string

const contextKey = key("app")

// FromContext get app from context
func FromContext(ctx context.Context) *App {
	app, ok := ctx.Value(contextKey).(*App)
	if !ok {
		return nil
	}
	return app
}

// ToContext put app to context
func ToContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, contextKey, app)
}

// Injector is middleware to inject app to request context
func Injector(app *App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ToContext(r.Context(), app)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
