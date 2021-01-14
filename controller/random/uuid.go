package random

import (
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/vetusbs/utils-api/controller/render"
)

func RandomUuid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := response{Value: guuid.New().String()}
		render.Render(w, response, http.StatusOK, r)
	}
}
