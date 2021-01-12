package controller

import (
	"github.com/gorilla/mux"
)

// Register is the main function of the http server
func Register() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/random/int", randomInt()).Methods("GET")
	api.HandleFunc("/random/string", randomString()).Methods("GET")

	return api
}
