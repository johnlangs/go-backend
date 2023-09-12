package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"	
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

var Logger TransactionLogger

var Config toml.Tree

func LoadConfig() string {

	Config, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	err = InitializeSQLTransactionLog(
		Config.Get("dbName").(string), 
		Config.Get("host").(string), 
		Config.Get("user").(string), 
		Config.Get("password").(string))
	if err != nil {
		panic(err)
	}

	return Config.Get("port").(string)
}


func KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
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

	Logger.WritePut(key, string(value))
	w.WriteHeader(http.StatusCreated)
}

func KeyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	Logger.WriteDelete(key)
	w.WriteHeader(http.StatusOK)
}

func KeyValueGetHandler(w http.ResponseWriter, r *http.Request) {
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}


func main() {
	time.Sleep(time.Millisecond * 5000)

	r := mux.NewRouter()

	port := LoadConfig()

	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/v1/{key}", KeyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", KeyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", KeyValueDeleteHandler).Methods("DELETE")

	
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "application/x-www-form-urlencoded", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})	

	
	println("Service Started. Listening on " + port)
	if port == ":443" {
		log.Fatal(http.ListenAndServeTLS(port, "cert.pem", "key.pem", r))
	} else {
		log.Fatal(http.ListenAndServe(port, r), handlers.CORS(originsOk, headersOk, methodsOk)(r))
	}
}