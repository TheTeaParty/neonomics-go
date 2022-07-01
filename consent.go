package neonomics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *client) GetConsent(ctx context.Context, ID string) (*GetConsentResponse, error) {

	_, err := c.TokenRequest(ctx)
	if err != nil {
		return nil, err
	}

	deviceID, ok := ctx.Value(CtxKeyDeviceID).(string)
	if !ok {
		return nil, ErrMissingCtxKey
	}

	headers := map[string]string{
		CtxKeyDeviceID:     deviceID,
		"accept":           "application/json",
		"authorization":    fmt.Sprintf("Bearer %s", c.tokenResponse.AccessToken),
		"content-type":     "application/json",
	}

	psuID, ok := ctx.Value(CtxKeyPSUID).(string)
	if ok {
		headers[CtxKeyPSUID] = psuID
	}

	psuIpAddress, ok := ctx.Value(CtxKeyPSUIPAddress).(string)
	if ok {
		headers[CtxKeyPSUIPAddress] = psuIpAddress
	}

	redirectURL, ok := ctx.Value(CtxKeyRedirectURL).(string)
	if ok {
		headers[CtxKeyRedirectURL] = redirectURL
	}

	reqURL := fmt.Sprintf(string(PathGetConsent), ID)

	responseBody, err := c.doRequest(ctx, reqURL, http.MethodGet, headers, nil)
	if err != nil {
		return nil, err
	}

	var response *GetConsentResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, err

}
