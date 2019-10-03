package postmark

import (
	"bytes"
	"encoding/json"
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
		return nil, err
	}

	// add headers for the postmark api
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", service.APIKey)

	// perform the request
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	// parse the results
	response := &Response{}
	json.NewDecoder(resp.Body).Decode(response)
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
