package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func render(w http.ResponseWriter, v interface{}, status int, r *http.Request) {

	accepts := r.Header.Get("Accept")
	if accepts == "text/plain" {
		writeText(w, v, status)
	} else {
		writeJSON(w, v, status)
	}

}

func writeText(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)

	io.WriteString(w, fmt.Sprintf("%v", v))
}

func writeJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}
