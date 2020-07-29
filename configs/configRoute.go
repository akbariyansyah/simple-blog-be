package configs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateRouter -> create new router
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
// StartServer -> start the server 
func StartServer(router *mux.Router) {
	host := ReadEnv("host")
	port := ReadEnv("port")
	server := fmt.Sprintf("%s:%s",host,port)
	log.Printf("Starting server at port : %s",port)
	err := http.ListenAndServe(server,router)

	if err != nil {
		log.Println(err)
	}
}