package utils

import (
	"encoding/json"
	baseResponse "lexicon/bo-api/common/models"
	"net/http"
)

// WriteMessage writes a JSON response with a custom message and status code.
// It takes an http.ResponseWriter, an integer status code, and a string message as parameters.
func WriteMessage(w http.ResponseWriter, status int, msg string) {
	var j struct {
		Msg string `json:"message"`
	}

	j.Msg = msg

	WriteResponse(w, j, status)
}

// WriteResponse writes a JSON response with the provided content and status code.
// It takes an http.ResponseWriter, an interface{} content, and an integer status code as parameters.
func WriteResponse(w http.ResponseWriter, content interface{}, status int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(content)
}

// WriteData writes a JSON response with the provided content as the data field and status code.
// It takes an http.ResponseWriter, an interface{} content, and an integer status code as parameters.
func WriteData(w http.ResponseWriter, content interface{}, status int) {
	var j = baseResponse.BaseResponse{
		Data: content,
	}

	WriteResponse(w, j, status)
}

// WriteError writes a JSON response with the provided error message, error status, and status code.
// It takes an http.ResponseWriter, an integer status code, and an error as parameters.
func WriteError(w http.ResponseWriter, status int, err error) {
	var j = baseResponse.ErrorResponse{
		Msg:   err.Error(),
		Error: http.StatusText(status),
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(j)
}
