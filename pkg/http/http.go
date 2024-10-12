package http

import (
	"encoding/json"
	"fmt"
	"io"
	netHttp "net/http"
)

// ClientImp defines the interface for the HTTP client implementation.
type ClientImp interface {
	FetchPublicData(url string, data interface{}) error
}

// httpHandler defines signature methods to mock "net/http" allowing futures tests.
type httpHandler interface {
	Get(url string) (resp *netHttp.Response, err error)
}

// client struct holds the instance of the HTTP client used for making requests.
type client struct {
	api httpHandler
}

// NewClient creates and returns a new instance of the HTTP client.
// It takes an *http.Client as an argument, allowing for customization of the HTTP client settings (e.g., timeouts, transport).
func NewClient(api httpHandler) ClientImp {
	return &client{api}
}

// FetchPublicData executes a GET request to the specified external public API URL.
// It decodes the JSON response into the provided data interface{} and handles potential errors during the request.
func (c *client) FetchPublicData(url string, data interface{}) error {
	response, err := c.api.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			fmt.Printf("Error closing response body: %s\n", closeErr.Error())
		}
	}(response.Body)

	return json.NewDecoder(response.Body).Decode(data)
}
