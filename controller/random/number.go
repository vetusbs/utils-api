package random

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/vetusbs/utils-api/controller/render"
)

func RandomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		max, _ := strconv.Atoi(r.URL.Query().Get("max"))

		if max == 0 {
			max = 1000
		}
		response := []interface{}{strconv.Itoa(rand.Intn(max))}
		render.Render(w, response, http.StatusOK, r)
	}
}
