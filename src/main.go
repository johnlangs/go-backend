package main

import (
	"log"
	"net/http"
	"go-backend/logging"
	"go-backend/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	logging.InitializeFileTransactionLog()

	r.HandleFunc("/", handlers.HelloMuxHandler)

	r.HandleFunc("/v1/{key}", handlers.KeyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", handlers.KeyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", handlers.KeyValueDeleteHandler).Methods("DELETE")

	println("Service Started. Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	//log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
}
