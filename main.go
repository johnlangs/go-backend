package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r* http.Request) {
	w.Write([]byte("Hello gorilla/mux\n"))
}

func ProductHandler(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(vars["key"]))
}

func ArticlesCategoryHandler(w http.ResponseWriter, r* http.Request) {

}

func ArticleHandler(w http.ResponseWriter, r* http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)
	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	println("Service Started. Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}