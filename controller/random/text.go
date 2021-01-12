package random

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/vetusbs/utils-api/controller/render"
)

func RandomString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		len, _ := strconv.Atoi(r.URL.Query().Get("len"))

		if len == 0 {
			len = 10
		}

		response := response{Value: randStringRunes(len)}

		render.Render(w, response, http.StatusOK, r)
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
