package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// sendMessageThroughPostmark will perform the sending operation
func (service *Service) sendMessageThroughPostmark(content *bytes.Buffer, path string) (*Response, error) {

	// create an endpoint url (with https and host address)
	url := makeEndpoint(service.Host, path)

	// create a new request
	req, err := http.NewRequest("POST", url.String(), content)
	if err != nil {
		return nil, fmt.Errorf("could not initialize post request: %w", err)
	}

	// add headers for the postmark api
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", service.APIKey)

	// perform the request
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform post request: %w", err)
	}

	// parse the results
	response := &Response{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, fmt.Errorf("could not decode server response: %w", err)
	}
	return response, nil
}

func makeEndpoint(host, path string) *url.URL {
	url := &url.URL{}

	url.Scheme = "https"

	if host == "" {
		url.Host = DefaultHost
	} else {
		url.Host = host
	}

	url.Path = path

	return url
}
