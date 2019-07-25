package apiserver

import (
	"time"

	"github.com/equest/exam/internal/api-server/handlers/auth"
	mauth "github.com/equest/exam/internal/api-server/middlewares/auth"

	"github.com/equest/exam/internal/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func buildRoutes(a *app.App) chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// A good base middleware stack
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(app.Injector(a))
	r.Use(mauth.Middleware(mauth.NewCognitoValidator(a.Services.Auths)))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chimw.Timeout(60 * time.Second))

	// prometheus handler
	r.Handle("/metrics", promhttp.Handler())

	// callback method

	// Public API
	r.Route("/v1", func(r chi.Router) {
		r.Get("/validate", auth.Validate())
	})

	return r
}
