package handlers

import (
	"fmt"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)// trail slashes will be ignored in here
    
	r.Route("/",func(router chi.Router) {
        fmt.Println("server")
		router.Get("/books/{id}",GetBook)
		router.Get("/books",GetAllBooks)
		router.Post("/books",SaveBook)
		router.Put("/books/{id}",UpdateBook)
		router.Delete("/books/{id}",DeleteBook)
		router.Get("/books/search",GetSearchBook)
	})
}

// added testing api sites for the purose good
