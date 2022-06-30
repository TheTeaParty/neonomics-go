package neonomics

import (
	"net/http"
)

type Config struct {
	ClientID string
	SecretID string
}

type Backend struct {
	Endpoint    Endpoint
	Environment Environment
	HttpClient  *http.Client
}
