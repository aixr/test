package main

import (
	"log"
	"net/http"
	
	"github.com/aixr/test/api"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
	mount(r, "/", api.router())
    log.Fatal(http.ListenAndServe(":8080", r))
}

func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}