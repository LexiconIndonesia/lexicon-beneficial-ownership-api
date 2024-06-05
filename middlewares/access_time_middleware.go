package middlewares

import (
	"net/http"
	"strconv"
	"time"
)

func AccessTime() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			now := time.Now()
			safeTime := now.Add(time.Minute * 3)

			reqTime := r.Header.Get("X-ACCESS-TIME")

			access, err := strconv.ParseFloat(reqTime, 64)

			if err != nil {
				middlewareError(w, http.StatusBadRequest, "Bad Request", "Invalid X-ACCESS-TIME Header")
				return
			}

			if access > float64(safeTime.Unix()) {
				middlewareError(w, http.StatusBadRequest, "Bad Request", "Invalid X-ACCESS-TIME Header")
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
