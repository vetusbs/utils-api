package random

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/vetusbs/utils-api/controller/render"
)

// ShowAccount godoc
// @Summary Generate random number
// @Description Generate random number from 0 to max. default value for max is 1000
// @ID generate-random-number
// @Accept  */*
// @Produce  text/plain
// @Param max path int true "Max Value"
// @Success 200 {object} render.Response
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} object
// @Failure 500 {object} object
// @Failure default {object} object
// @Router /random/int [get]
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
