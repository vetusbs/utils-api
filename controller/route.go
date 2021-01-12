package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Register is the main function of the http server
func Register() *mux.Router {
	api := mux.NewRouter()

	api.Use(commonMiddleware)

	api.HandleFunc("/random/int", randomInt()).Methods("GET")
	api.HandleFunc("/random/string", randomString()).Methods("GET")

	return api
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}
