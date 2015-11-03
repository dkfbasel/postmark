// Package postmark provides a wrapper for the postmarkapi.com email service
package postmark

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	// DefaultHost is used to hold the standard config to send messages to
	DefaultHost = "api.postmarkapp.com"
)

// Service contains all information and methods for the postmark api service
// Host will be set to the default host "api.postmarkapp.com" if not specified
type Service struct {
	APIKey string
	Host   string
}

// Response will contain the information returned by the postmark-api
type Response struct {
	ErrorCode   int
	Message     string
	MessageID   string
	SubmittedAt string
	To          string
}

// Send swill send a single message to the server
func (service *Service) Send(msg *Message) (*Response, error) {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, err
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email")
}

// SendWithTemplate will send a message using a pre-specified template
func (service *Service) SendWithTemplate(msg *MessageWithTemplate) (*Response, error) {

	// marshal the message
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, err
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email/withTemplate")
}

// SendBatch will send multiple messages using the batch API
func (service *Service) SendBatch(msg []*Message) (*Response, error) {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, err
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email/batch")
}

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
