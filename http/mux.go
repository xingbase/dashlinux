package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type MuxOpts struct {
	PprofEnabled bool // Mount pprof routes for profiling
}

func NewMux(opts MuxOpts, service Service) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if opts.PprofEnabled {
		r.Mount("/debug", middleware.Profiler())
	}

	r.Get("/ping", service.Ping)

	r.Post("/auth/login", service.Login)
	r.Post("/auth/token", service.Login)

	// Ensure routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator)

		r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})
	return r
}
