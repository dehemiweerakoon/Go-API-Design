package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
)

func GetSearchBook(w http.ResponseWriter,r *http.Request){
	//fmt.Println(r.)
    var params = api.SearchQuery{}
	var decoder *schema.Decoder = schema.NewDecoder()
    var err error
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		fmt.Println("hi")
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database,err = tools.NewDatabase()

	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}
	var bookDetails = (*database).SearchHandler(params.Query)
	if bookDetails==nil{
		fmt.Println(bookDetails)
		log.Error(err)
		res := api.SearchQuery{
			Query: "No Books Found",
		}
		w.Header().Set("Content-Type","application/json")
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
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