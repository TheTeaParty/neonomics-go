package neonomics

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type client struct {
	config  *Config
	backend *Backend
}

func (c *client) doRequest(ctx context.Context, path Path, method string,
	headers map[string]string, requestBody []byte) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, method,
		fmt.Sprintf("%s%s", c.backend.Endpoint, path), bytes.NewReader(requestBody))
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
