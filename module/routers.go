package module

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {

	r := chi.NewMux()
	r.Get("/", testRoute)
	return r
}

func writeMessage(w http.ResponseWriter, status int, msg string) {
	var j struct {
		Msg string `json:"message"`
	}

	j.Msg = msg

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)
}

func writeData(w http.ResponseWriter, content interface{}, status int) {
	var j struct {
		Data interface{} `json:"data"`
	}

	j.Data = content
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeMessage(w, status, err.Error())
}

func testRoute(w http.ResponseWriter, req *http.Request) {

	writeMessage(w, 200, "Hello")
}
