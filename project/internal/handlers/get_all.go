package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
)

func GetAllBooks(w http.ResponseWriter,r *http.Request){

	var err error

	var database *tools.DatabaseInterface
	database,err = tools.NewDatabase()

	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}
	var bookDetails = (*database).GetAllBookDetails()
	if bookDetails==nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var responses []api.BookResponseParam
	for _, book := range bookDetails {
		response := api.BookResponseParam{
			BookId: book.BookId,
			AuthorId: book.AuthorId,
			PublisherId: book.PublisherId,
			Title: book.Title,
			PublicationDate: book.PublicationDate,
			Isbn: book.Isbn,
			Pages: book.Pages,
			Genre: book.Genre,
			Description: book.Description,
			Price: book.Price,
			Quantity: book.Quantity,
		}
		responses = append(responses, response)
	}
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(responses)
	if err!=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}