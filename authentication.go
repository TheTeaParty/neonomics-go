package neonomics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *client) TokenRequest(ctx context.Context) (*TokenRequestResponse, error) {

	req := &TokenRequestRequest{
		GrantType:    "client_credentials",
		ClientId:     c.config.ClientID,
		ClientSecret: c.config.SecretID,
	}

	data := url.Values{}
	data.Set("grant_type", req.GrantType)
	data.Set("client_id", req.ClientId)
	data.Set("client_secret", req.ClientSecret)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	reqURL := fmt.Sprintf(string(PathTokenRequest), "sandbox")
	if c.backend.Environment == EnvironmentProduction {
		reqURL = fmt.Sprintf(string(PathTokenRequest), "live")
	}

	responseBody, err := c.doRequest(ctx, reqURL, http.MethodPost, headers, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var response TokenRequestResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) TokenRefresh(ctx context.Context) (*TokenRefreshResponse, error) {
	//TODO implement me
	panic("implement me")
}
