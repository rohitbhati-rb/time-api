package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// our custom multiplexer from gorilla/mux
	router := mux.NewRouter()

	// time route
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}
