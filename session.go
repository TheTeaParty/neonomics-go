package neonomics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *client) CreateSession(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {

	_, err := c.TokenRequest(ctx)
	if err != nil {
		return nil, err
	}

	deviceID, ok := ctx.Value(CtxKeyDeviceID).(string)
	if !ok {
		return nil, ErrMissingCtxKey
	}

	headers := map[string]string{
		CtxKeyDeviceID:  deviceID,
		"accept":        "application/json",
		"authorization": fmt.Sprintf("Bearer %s", c.tokenResponse.AccessToken),
		"content-type":  "application/json",
	}

	reqURL := fmt.Sprintf(string(PathCreateSession))

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := c.doRequest(ctx, reqURL, http.MethodPost, headers, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	var response *CreateSessionResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, err
}

func (c *client) GetSessionStatus(ctx context.Context, ID string) (*GetSessionStatusResponse, error) {

	_, err := c.TokenRequest(ctx)
	if err != nil {
		return nil, err
	}

	deviceID, ok := ctx.Value(CtxKeyDeviceID).(string)
	if !ok {
		return nil, ErrMissingCtxKey
	}

	headers := map[string]string{
		CtxKeyDeviceID:  deviceID,
		"accept":        "application/json",
		"authorization": fmt.Sprintf("Bearer %s", c.tokenResponse.AccessToken),
		"content-type":  "application/json",
	}

	reqURL := fmt.Sprintf(string(PathGetSessionStatus), ID)

	responseBody, err := c.doRequest(ctx, reqURL, http.MethodGet, headers, nil)
	if err != nil {
		return nil, err
	}

	var response *GetSessionStatusResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, err

}

func (c *client) DeleteSession(ctx context.Context, ID string) error {

	_, err := c.TokenRequest(ctx)
	if err != nil {
		return err
	}

	deviceID, ok := ctx.Value(CtxKeyDeviceID).(string)
	if !ok {
		return ErrMissingCtxKey
	}

	headers := map[string]string{
		CtxKeyDeviceID:  deviceID,
		"accept":        "application/json",
		"authorization": fmt.Sprintf("Bearer %s", c.tokenResponse.AccessToken),
		"content-type":  "application/json",
	}

	reqURL := fmt.Sprintf(string(PathDeleteSession), ID)

	_, err = c.doRequest(ctx, reqURL, http.MethodDelete, headers, nil)
	if err != nil {
		return err
	}

	return err
}
