package middlewares

import (
	"encoding/json"
	"net/http"
)

func middlewareError(w http.ResponseWriter, s int, err string, m string) {

	var j struct {
		Error string `json:"error"`
		Msg   string `json:"message"`
	}

	j.Error = err
	j.Msg = m

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(s)

	json.NewEncoder(w).Encode(j)

}
