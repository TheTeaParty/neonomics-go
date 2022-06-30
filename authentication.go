package neonomics

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *client) TokenRequest(ctx context.Context, req *TokenRequestRequest) (*TokenRequestResponse, error) {

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	responseBody, err := c.doRequest(ctx, PathTokenRequest, http.MethodPost, headers, requestBody)
	if err != nil {
		return nil, err
	}

	var response *TokenRequestResponse
	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) TokenRefresh(ctx context.Context, req *TokenRefreshRequest) (*TokenRefreshResponse, error) {
	//TODO implement me
	panic("implement me")
}
