package neonomics

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type client struct {
	config  *Config
	backend *Backend
}

func (c *client) doRequest(ctx context.Context, path string, method string,
	headers map[string]string, requestBody io.Reader) ([]byte, error) {

	reqURL := fmt.Sprintf("%s%s", c.backend.Endpoint, path)

	req, err := http.NewRequestWithContext(ctx, method, reqURL, requestBody)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	rsp, err := c.backend.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode >= 400 {
		var errResponse ErrorResponse
		err := json.Unmarshal(body, &errResponse)
		if err != nil {
			return nil, ErrUnexpectedError
		}

		return nil, errors.New(errResponse.Message)
	}

	return body, nil
}

func New(config *Config, backend *Backend) API {
	return &client{
		config:  config,
		backend: backend,
	}
}

func NewProduction(config *Config) API {
	backend := &Backend{
		Endpoint:    EndpointProduction,
		HttpClient:  &http.Client{},
		Environment: EnvironmentProduction,
	}

	return New(config, backend)
}

func NewSandbox(config *Config) API {
	backend := &Backend{
		Endpoint:    EndpointSandbox,
		HttpClient:  &http.Client{},
		Environment: EnvironmentSandbox,
	}

	return New(config, backend)
}
