package go_cod

import (
	"context"
	"fmt"
	"github.com/carlocayos/go-cod/v2/api/client/authentication"
	"github.com/carlocayos/go-cod/v2/api/models"
	"net/http"
)

// RegisterDevice registers the device and returns the Auth Header token
func (c *Client) RegisterDevice(ctx context.Context, deviceId string) (*models.RegisterDeviceResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := authentication.RegisterDeviceParams{
		RegisterDeviceRequest: &models.RegisterDeviceRequest{DeviceID: &deviceId},
		Context:               ctx,
	}

	res, err := c.AuthenticationClient.Operations.RegisterDevice(&param)
	if err != nil {
		return nil, err
	}

	return res.Payload, nil
}

// Login uses the token received from register device authHeader, then returns cookie security tokens
//	used for all authenticated requests
func (c *Client) Login(ctx context.Context, deviceId string, email string, password string, token string) (*models.LoginResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	l := authentication.LoginParams{
		Authorization: "Bearer " + token,
		XCodDeviceID:  deviceId,
		LoginRequest: &models.LoginRequest{
			Email:    &email,
			Password: &password,
		},
		Context: context.Background(),
	}

	res, err := c.AuthenticationClient.Operations.Login(&l)
	if err != nil {
		return nil, err
	}
	// For an unauthorized request, API returns HTTP code 200, but with payload status = 401
	if res.Payload.Status == http.StatusUnauthorized {
		return nil, fmt.Errorf("%v", res.Payload.Message)
	}

	// store the Activision sso cookie token
	c.AuthToken.ActSSOCookie = res.Payload.ACTSSOCOOKIE

	return res.Payload, nil
}
