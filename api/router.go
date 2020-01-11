package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", get).Methods(http.MethodGet)
	return r
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}