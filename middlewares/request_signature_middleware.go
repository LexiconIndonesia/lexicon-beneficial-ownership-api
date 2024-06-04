package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func RequestSignature(salt string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if len(salt) <= 0 {
				middlewareError(w, http.StatusInternalServerError, "Internal Server Error", "No Salt Found")
				return
			}

			accessTime := r.Header.Get("X-ACCESS-TIME")
			apiKey := r.Header.Get("X-API-KEY")
			signature := r.Header.Get("X-REQUEST-SIGNATURE")

			if len(accessTime) <= 0 {
				middlewareError(w, http.StatusBadRequest, "Bad Request", "Missing X-ACCESS-TIME")
				return
			}

			if len(apiKey) <= 0 {
				middlewareError(w, http.StatusForbidden, "Forbidden", "Missing X-API-KEY")
				return
			}
			if len(signature) <= 0 {
				middlewareError(w, http.StatusForbidden, "Forbidden", "Missing X-REQUEST-SIGNATURE")
				return
			}
			hash := sha256.New()
			hash.Write([]byte(salt + accessTime + apiKey))

			hashedSignature := hex.EncodeToString(hash.Sum(nil))

			if signature != hashedSignature {
				middlewareError(w, http.StatusBadRequest, "Bad Request", "Invalid X-API-KEY Header")
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}
