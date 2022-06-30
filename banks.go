package neonomics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *client) GetSupportedBanks(ctx context.Context) ([]*GetSupportedBanksResponse, error) {

	deviceID, ok := ctx.Value(CtxKeyDeviceID).(string)
	if !ok {
		return nil, ErrMissingCtxKey
	}

	headers := map[string]string{
		CtxKeyDeviceID:  deviceID,
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer: %s", c.tokenResponse.AccessToken),
	}

	responseBody, err := c.doRequest(ctx, string(PathGetSupportedBanks), http.MethodGet, headers, nil)
	if err != nil {
		return nil, err
	}

	var response []*GetSupportedBanksResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, err
}

func (c *client) GetSupportedBankByID(ctx context.Context, ID string) (*GetSupportedBanksResponse, error) {
	//TODO implement me
	panic("implement me")
}
