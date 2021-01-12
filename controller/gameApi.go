package controller

import (
	"encoding/json"
	"net/http"
)

func test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			js, _ := json.Marshal(Response{Value: "test"})
			w.Write(js)
		}
	}
}

type Response struct {
	Value string `json:"value"`
}
