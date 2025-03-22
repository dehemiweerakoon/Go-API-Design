package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
)

func UpdateBook(w http.ResponseWriter,r *http.Request){
	id := chi.URLParam(r,"id")
	params_list := api.BookBalanceParam{BookId: id}
	var params = api.BookResponseParam{}

	decoder := json.NewDecoder(r.Body)

	if err:=decoder.Decode(&params);err!=nil{
		log.Error(err)
		api.RequestErrorHandler(w,err)
		return
	}
	if params.BookId == "" {
		log.Error("No Book Id provided for the book save ")
		//api.RequestErrorHandler(w,)
	}
	var err error

	var database *tools.DatabaseInterface
	database,err = tools.NewDatabase()

	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}
	var bookDetails = (*database).UpdateBookDetails(tools.BookDetails(params),params_list.BookId)
	if bookDetails==nil{
		log.Error(err)
		api.InternalErrorHandler(w)
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