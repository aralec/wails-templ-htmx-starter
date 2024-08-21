package render

import (
	"wails-templ-hmtx-project/src/internal/adapters/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// GetRenderingMux creates a new router for the application that serves views rendered by Go Templ
func GetRenderingMux() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/counter", func(r chi.Router) {
		cs := handlers.NewCounterHandler()
		r.Get("/", cs.ServeCounterView)
		r.Post("/increment", cs.Increment)
		r.Post("/decrement", cs.Decrement)
	})

	return r
}
