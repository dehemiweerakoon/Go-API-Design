package main

import (
	"fmt"
	"net/http"

	"github.com/dehemiweerakoon/golan-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go API Services")

	fmt.Println(`
####### ########   #            #     #   # ####### 
#       #      #   #           # #    ##  # #
#  #### #      #   #          #####   # # # #  ####
#     # #      #   #         #     #  #  ## #     #
####### ########   ######   #       # #   # #######
	`)
	err := http.ListenAndServe("localhost:9000",r)
	if err !=nil{
		log.Error(err)
	}
}