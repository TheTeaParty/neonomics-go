package neonomics

import (
	"net/http"
)

type client struct {
	config  *Config
	backend *Backend
}

func New(config *Config, backend *Backend) API {
	return &client{
		config:  config,
		backend: backend,
	}
}

func NewProduction(config *Config) API {
	backend := &Backend{
		Endpoint:   EndpointProduction,
		HttpClient: &http.Client{},
	}

	return New(config, backend)
}

func NewSandbox(config *Config) API {
	backend := &Backend{
		Endpoint:   EndpointSandbox,
		HttpClient: &http.Client{},
	}

	return New(config, backend)
}
