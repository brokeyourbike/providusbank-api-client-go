package providusbank

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	httpClient HttpClient
	logger     *logrus.Logger
}

// ClientOption is a function that configures a Client.
type ClientOption func(*client)

// WithHTTPClient sets the HTTP client for the paystack API client.
func WithHTTPClient(c HttpClient) ClientOption {
	return func(target *client) {
		target.httpClient = c
	}
}

// WithLogger sets the *logrus.Logger for the paystack API client.
func WithLogger(l *logrus.Logger) ClientOption {
	return func(target *client) {
		target.logger = l
	}
}
