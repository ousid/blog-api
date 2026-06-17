package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/coderflexx/blog-api/internal/handlers"
)

func Setup() http.Handler {
	r := chi.NewRouter()

	// Global Middleware
	r.Use(middleware.Logger)    // logs every request
	r.Use(middleware.Recoverer) // catches panics
	r.Use(middleware.CleanPath) // Normalizes URLs

	// health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	// API routes group
	r.Route("/api", func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", handlers.ListPosts)
			r.Get("/{slug}", handlers.GetPost)
			r.Post("/", handlers.CreatePost)
			r.Put("/{id}", handlers.UpdatePost)
			r.Delete("/{id}", handlers.DeletePost)
		})

		r.Route("/categories", func(r chi.Router) {
			//
		})
	})

	return r
}
