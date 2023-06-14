package main

import (
	"io"
	"log"
	"net/http"
	"errors"

	"github.com/gorilla/mux"
)

var store = make(map[string]string)

func Put(key string, value string) error {
	store[key] = value

	return nil
}

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)

	return nil
}

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

func keyValuePutHandler (w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
			return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusCreated)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)
	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	r.HandleFunc("/v1/{key}", keyValuePutHandler,).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHandler,).Methods("GET")

	println("Service Started. Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}