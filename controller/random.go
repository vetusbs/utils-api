package controller

import (
	"math/rand"
	"net/http"
	"strconv"
)

func randomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		max, _ := strconv.Atoi(r.URL.Query().Get("max"))

		if max == 0 {
			max = 1000
		}
		response := Response{Value: strconv.Itoa(rand.Intn(max))}
		render(w, response, http.StatusOK, r)
	}
}

func randomString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		len, _ := strconv.Atoi(r.URL.Query().Get("len"))

		if len == 0 {
			len = 10
		}

		response := Response{Value: randStringRunes(len)}

		render(w, response, http.StatusOK, r)
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
