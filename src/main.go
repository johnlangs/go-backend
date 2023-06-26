package main

import (
	"go-backend/utils"
	"go-backend/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	
	port := utils.LoadConfig()

	r.HandleFunc("/", handlers.HelloMuxHandler)

	r.HandleFunc("/v1/{key}", handlers.KeyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", handlers.KeyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", handlers.KeyValueDeleteHandler).Methods("DELETE")

	if port == ":443" {
		log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", r))
	} else {
		log.Fatal(http.ListenAndServe(port, r))
	}
	println("Service Started. Listening on " + port)
}
