package utils

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func LaunchRPC(handler http.Handler, port string) error {
	r := mux.NewRouter()
	r.Handle("/", handler).Methods("GET", "POST", "OPTIONS")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(r))
}
