package utils

import "net/http"

var (
	Client *http.Client
)

func SetClient(client *http.Client) {
	Client = client
}
