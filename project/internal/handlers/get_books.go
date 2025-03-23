package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
)

func GetBook(w http.ResponseWriter,r *http.Request){
	// var params = api.BookBalanceParam{}
	// var decoder  *schema.Decoder = schema.NewDecoder()
    // err = decoder.Decode(&params,r.URL.Query())
	var err error
	

	 // Extract path variable "id" using chi.URLParam
	 id := chi.URLParam(r, "id")

	 // Map the "id" to the struct or use it directly
	 params := api.BookBalanceParam{BookId: id}

	var database *tools.DatabaseInterface
	database,err = tools.NewDatabase()

	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}
	var bookDetails = (*database).GetBookDetails(params.BookId)
	if bookDetails==nil{
		var response = api.BookResponseParam{
			BookId: "No book Is found",
		}
		w.Header().Set("Content-Type","application/json")
		err = json.NewEncoder(w).Encode(response)
		if err!=nil{
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
		return
	}
	var response = api.BookResponseParam{
		BookId: (*bookDetails).BookId,
		AuthorId: (*bookDetails).AuthorId,
		PublisherId: (*bookDetails).PublisherId,
		Title: (*bookDetails).Title,
		PublicationDate: (*bookDetails).PublicationDate,
		Isbn: (*bookDetails).Isbn,
		Pages: (*bookDetails).Pages,
		Genre: (*bookDetails).Genre,
		Description: (*bookDetails).Description,
		Price: (*bookDetails).Price,
		Quantity: (*bookDetails).Quantity,
	}
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(response)
	if err!=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}