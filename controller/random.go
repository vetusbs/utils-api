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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		max, _ := strconv.Atoi(r.URL.Query().Get("max"))

		if max == 0 {
			max = 1000
		}

		js, _ := json.Marshal(Response{Value: strconv.Itoa(rand.Intn(max))})
		w.Write(js)

	}
}

func randomString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		len, _ := strconv.Atoi(r.URL.Query().Get("len"))

		if len == 0 {
			len = 10
		}

		js, _ := json.Marshal(Response{Value: randStringRunes(len)})
		w.Write(js)

	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Response struct {
	Value string `json:"value"`
}
