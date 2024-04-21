package utils

import (
	"encoding/json"
	"net/http"
)

func WriteMessage(w http.ResponseWriter, status int, msg string) {
	var j struct {
		Msg string `json:"message"`
	}

	j.Msg = msg

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)
}

func WriteResponse(w http.ResponseWriter, content interface{}) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(content)
}

func WriteData(w http.ResponseWriter, content interface{}, status int) {
	var j struct {
		Data interface{} `json:"data"`
	}

	j.Data = content

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	var j struct {
		Error string `json:"error"`
		Msg   string `json:"message"`
	}

	j.Msg = err.Error()
	j.Error = http.StatusText(status)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)

}
