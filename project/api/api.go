package api

import (
	"encoding/json"
	"net/http"
)
type DeleteElement struct{
	DeletedOrNot string
}

type BookBalanceParam struct {
	BookId string
}
type UserParam struct {
	Username string
}

type BookResponseParam struct {
	BookId          string
	AuthorId        string
	PublisherId     string
	Title           string
	PublicationDate string
	Isbn            string
	Pages           int
	Genre           string
	Description     string
	Price float64
	Quantity int 
}
type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter,message string,code int) {
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func (w http.ResponseWriter,err error)  {
		writeError(w,err.Error(),http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w,"An Unexpected Error Occurred",http.StatusInternalServerError)

	}
)