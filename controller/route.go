package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vetusbs/utils-api/controller/random"
)

// Register is the main function of the http server
func Register() *mux.Router {
	api := mux.NewRouter()

	api.Use(commonMiddleware)

	api.HandleFunc("/random/int", random.RandomInt()).Methods("GET")
	api.HandleFunc("/random/string", random.RandomString()).Methods("GET")

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
