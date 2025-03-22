package handlers

import (
	"github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
	//"github.com/dehemiweerakoon/golan-api/internal/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)// trail slashes will be ignored in here
    
	r.Route("/",func(router chi.Router) {
		//router.Use(middleware.Authorization)

		router.Get("/books/{id}",GetBook)
		router.Get("/books",GetAllBooks)
		router.Post("/books",SaveBook)
	})
}
