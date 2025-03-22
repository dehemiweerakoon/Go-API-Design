package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dehemiweerakoon/golan-api/api"
	"github.com/dehemiweerakoon/golan-api/internal/tools"
	log "github.com/sirupsen/logrus"
)
var ErrUnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var username string = r.URL.Query().Get("username")
		var params  = api.UserParam{}
		decoder :=json.NewDecoder(r.Body)
		if err := decoder.Decode(&params); err != nil {
			log.Error(err)
			api.RequestErrorHandler(w, err)
			return
		}
		var token = r.Header.Get("Authorization")

		var err error

		if params.Username == "" || token == "" {
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w,ErrUnAuthorized)
		}
		var database *tools.DatabaseInterface
		database ,err = tools.NewDatabase()

		if err!=nil{
			api.InternalErrorHandler(w)
			return
		}
		loginDetails := (*database).GetUserLoginDetails(params.Username)
		if(loginDetails ==nil || (token!=((loginDetails).AuthToken))){
			log.Error(ErrUnAuthorized)
			api.RequestErrorHandler(w,ErrUnAuthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}
