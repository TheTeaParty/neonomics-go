package neonomics

import (
	"context"
	"reflect"
	"testing"
)

func Test_client_TokenRefresh(t *testing.T) {
	type fields struct {
		config  *Config
		backend *Backend
	}
	type args struct {
		ctx context.Context
		req *TokenRefreshRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TokenRefreshResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				config:  tt.fields.config,
				backend: tt.fields.backend,
			}
			got, err := c.TokenRefresh(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenRefresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenRefresh() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_TokenRequest(t *testing.T) {
	type fields struct {
		config  *Config
		backend *Backend
	}
	type args struct {
		ctx context.Context
		req *TokenRequestRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TokenRequestResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				config:  tt.fields.config,
				backend: tt.fields.backend,
			}
			got, err := c.TokenRequest(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
