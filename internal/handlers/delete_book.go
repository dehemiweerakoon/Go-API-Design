package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
)

func DeleteBook(w http.ResponseWriter,r *http.Request){

	var err error
	// err = decoder.Decode(&params,r.URL.Query())

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
	var bookDetails = (*database).DeleteBookDetails(params.BookId)
	if !bookDetails {
		log.Error(err)
		var response = api.DeleteElement{
			DeletedOrNot: "Book data is Not FOund ",
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
	var response = api.DeleteElement{
		DeletedOrNot: "Book data is Deleted ",
	}
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(response)
	if err!=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}