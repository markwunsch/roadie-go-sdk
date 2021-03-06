package roadie

import (
	"context"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
)

const (
	// DefaultRoadieHost is the default host address to use for calls to the roadie api
	DefaultRoadieHost = `https://connect.roadie.com`
	// DefaultRoadieVersion is the default version to use for calls to the roadie api
	DefaultRoadieVersion = `v1`
)

type service struct {
	client *Client
}

// Client is used to perform all operations with the roadie api
type Client struct {
	// host holds the address of the roadie server
	host string
	// client handles http requests
	client *http.Client
	// version of the roadie api to use
	version string
	// customHTTPHeaders can be set by user to include additional headers
	customHTTPHeaders map[string]string
	// customVersion is set to true if the user specified a custom version
	customVersion bool
	// accessToken is the user's access token to use for authorization
	accessToken string
	// Estimates is the service used to create estimates
	Estimates *EstimateService
	// Shipments is the service used to interact with shipments
	Shipments *ShipmentsService
}

// NewClient creates a new instance of Client with any optional functions applied
func NewClient(optFns ...func(*Client) error) (*Client, error) {
	client, err := defaultHTTPClient()
	if err != nil {
		return nil, err
	}
	c := &Client{
		host:    DefaultRoadieHost,
		version: DefaultRoadieVersion,
		client:  client,
	}
	c.Estimates = &EstimateService{client: c}
	c.Shipments = &ShipmentsService{client: c}

	for _, fn := range optFns {
		if err := fn(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithAccessToken allows the user to provide an access token to use for authorization
func WithAccessToken(ctx context.Context, accessToken string) func(*Client) error {
	return func(c *Client) error {
		c.accessToken = accessToken
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		tc.Timeout = time.Minute
		c.client = tc
		return nil
	}
}

// UpdateAccessToken can be used to update an outdated access token
func (cli *Client) UpdateAccessToken(ctx context.Context, accessToken string) {
	cli.accessToken = accessToken
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	cli.client = tc
}

// WithHTTPClient allows user to provide a custom client
func WithHTTPClient(client *http.Client) func(*Client) error {
	return func(c *Client) error {
		if client != nil {
			c.client = client
		}
		return nil
	}
}

// WithHost allows the user to use a custom roadie api host
func WithHost(host string) func(*Client) error {
	return func(c *Client) error {
		c.host = host
		return nil
	}
}

// WithVersion allows the user to use a custom roadie api version
func WithVersion(version string) func(*Client) error {
	return func(c *Client) error {
		c.version = version
		c.customVersion = true
		return nil
	}
}

// WithHTTPHeaders allows the user to specify custom headers to be used with all requests
func WithHTTPHeaders(headers map[string]string) func(*Client) error {
	return func(c *Client) error {
		c.customHTTPHeaders = headers
		return nil
	}
}

// GetCustomHTTPHeaders returns custom http headers stored by the client
func (cli *Client) GetCustomHTTPHeaders() map[string]string {
	headers := make(map[string]string)
	for k, v := range cli.customHTTPHeaders {
		headers[k] = v
	}
	return headers
}

// WithEnvVars allows retrieves host/version from environment variables
func WithEnvVars(c *Client) error {
	if host := os.Getenv("TWITTER_ADS_HOST"); host != "" {
		if err := WithHost(host)(c); err != nil {
			return err
		}
	}
	if version := os.Getenv("TWITTER_ADS_API_VERSION"); version != "" {
		if err := WithVersion(version)(c); err != nil {
			return err
		}
	}
	return nil
}

// defaultHTTPClient returns an http client with default parameters
func defaultHTTPClient() (*http.Client, error) {
	return &http.Client{
		Timeout: time.Minute,
	}, nil
}
