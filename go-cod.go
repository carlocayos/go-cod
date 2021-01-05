package go_cod

import (
	"github.com/carlocayos/go-cod/api/client/authentication"
	"github.com/carlocayos/go-cod/api/client/service"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	DefaultServiceHost            string = "my.callofduty.com"
	DefaultServiceBasePath        string = "/api/papi-client"
	DefaultAuthenticationHost     string = "profile.callofduty.com"
	DefaultAuthenticationBasePath string = "/"
)

var DefaultSchemes = []string{"https"}

// AuthenticationClient is the client containing authentication operations
type AuthenticationClient struct {
	Operations authentication.ClientService
	Transport  runtime.ClientTransport
}

// ServiceClient is the client containing common service operations
type ServiceClient struct {
	Operations service.ClientService
	Transport  runtime.ClientTransport
}

// Go COD Client contains both an Authentication and Service Client
//   1. Authentication client is usedfor authentication
//   2. Service client is for common api requests
type Client struct {
	AuthenticationClient *AuthenticationClient
	ServiceClient        *ServiceClient
}

// Creates a default COD Client
func NewClient(formats strfmt.Registry) *Client {
	// create the Go COD Authentication Client
	authClient := NewAuthenticationClientWithConfig(formats, DefaultAuthTransportConfig())

	// create the Go COD Service Client
	serviceClient := NewServiceClientWithConfig(formats, DefaultServiceTransportConfig())

	// create Client struct
	c := &Client{
		AuthenticationClient: authClient,
		ServiceClient:        serviceClient,
	}
	return c
}

// NewAuthenticationClientWithConfig creates a new Authentication client with config
func NewAuthenticationClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *AuthenticationClient {
	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return NewAuthenticationClient(transport, formats)
}

// NewServiceClientWithConfig creates a new Service client with config
func NewServiceClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ServiceClient {
	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return NewServiceClient(transport, formats)
}

// New creates a new Authentication client
func NewAuthenticationClient(transport runtime.ClientTransport, formats strfmt.Registry) *AuthenticationClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(AuthenticationClient)
	cli.Transport = transport
	cli.Operations = authentication.New(transport, formats)
	return cli
}

// New creates a new Service client
func NewServiceClient(transport runtime.ClientTransport, formats strfmt.Registry) *ServiceClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ServiceClient)
	cli.Transport = transport
	cli.Operations = service.New(transport, formats)
	return cli
}

// DefaultAuthTransportConfig creates a TransportConfig with the
// default settings for the Authentication Client
func DefaultAuthTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultAuthenticationHost,
		BasePath: DefaultAuthenticationBasePath,
		Schemes:  DefaultSchemes,
	}
}

// DefaultServiceTransportConfig creates a TransportConfig with the
// default settings for the Service Client
func DefaultServiceTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultServiceHost,
		BasePath: DefaultServiceBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SetTransport changes the transport on the Authentication client and all its subresources
func (c *AuthenticationClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Operations.SetTransport(transport)
}

// SetTransport changes the transport on the Service client and all its subresources
func (c *ServiceClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Operations.SetTransport(transport)
}
