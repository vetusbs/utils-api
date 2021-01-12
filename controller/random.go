package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

func randomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			js, _ := json.Marshal(Response{Value: strconv.Itoa(rand.Intn(1000))})
			w.Write(js)
		}
	}
}

type Response struct {
	Value string `json:"value"`
}
