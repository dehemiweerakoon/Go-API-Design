package handlers 

import (
	"github.com/go-chi/chi"
	chimiddle"github.com/go-chi/chi/middleware"
	"github.com/dehemiweerakoon/golan-api/internal/middleware"
)
func Handler(r *chi.Mux){
	// global middleware
	r.Use(chimiddle.StripSlashes) // trail slashes will be ignored 

	r.Route("/account",func(router chi.Router) {
		router.Use(middleware.Authorization)
		// Middleware for /account route

		router.Get("/coins",GetCoinBalance)
	})
}