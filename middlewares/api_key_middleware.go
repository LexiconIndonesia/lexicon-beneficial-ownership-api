package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func ApiKey(serverApiKeys string, salt string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if len(serverApiKeys) <= 0 || len(salt) <= 0 {
				middlewareError(w, http.StatusInternalServerError, "Internal Server Error", "No API Key Found")
				return
			}

			hostname := r.Header.Get("X-REQUEST-IDENTITY")
			apiKey := r.Header.Get("X-API-KEY")

			if len(hostname) <= 0 {
				middlewareError(w, http.StatusForbidden, "Forbidden", "Missing X-REQUEST-IDENTITY")
				return
			}

			if len(apiKey) <= 0 {
				middlewareError(w, http.StatusForbidden, "Forbidden", "Missing X-API-KEY")
				return
			}

			hash := sha256.New()
			hash.Write([]byte(salt + apiKey))

			hashedKey := hex.EncodeToString(hash.Sum(nil))

			accessedKey := serverApiKeys

			if accessedKey == "" {
				middlewareError(w, http.StatusForbidden, "Forbidden", "Invalid X-API-KEY Header")
				return
			}

			if accessedKey != hashedKey {
				middlewareError(w, http.StatusBadRequest, "Bad Request", "Invalid X-API-KEY Header")
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}
