// Package postmark provides a wrapper for the postmarkapi.com email service
package postmark

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// DefaultHost is used to hold the standard config to send messages to
	DefaultHost = "api.postmarkapp.com"
	TestApiKey  = "POSTMARK_API_TEST"
)

// Service contains all information and methods for the postmark api service
// Host will be set to the default host "api.postmarkapp.com" if not specified
type Service struct {
	APIKey string
	Host   string
	From   string
}

// Response will contain the information returned by the postmark-api
type Response struct {
	ErrorCode   int
	Message     string
	MessageID   string
	SubmittedAt string
	To          string
}

// New will return an email service instance that can be used to send emails
// via the service https://postmarkapp.com
func New(postmarkHost string, postmarkApikey string, defaultFrom string) (*Service, error) {

	if postmarkHost == "" || postmarkApikey == "" || defaultFrom == "" {
		return nil, errors.New("Postmark host, postmark api key and default sender address must be provided")
	}

	return &Service{
		APIKey: postmarkApikey,
		Host:   postmarkHost,
		From:   defaultFrom,
	}, nil

}

// Send swill send a single message to the server
func (service *Service) Send(msg *Message) (*Response, error) {

	// set default sender if not specified
	if msg.From == "" {
		msg.From = service.From
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, fmt.Errorf("could not encode email message: %w", err)
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email")
}

// SendWithTemplate will send a message using a pre-specified template
func (service *Service) SendWithTemplate(msg *MessageWithTemplate) (*Response, error) {

	// set default sender if not specified
	if msg.From == "" {
		msg.From = service.From
	}

	// marshal the message
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, fmt.Errorf("could not encode email message with template: %w", err)
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email/withTemplate")
}

// SendBatch will send multiple messages using the batch API
func (service *Service) SendBatch(msg []*Message) (*Response, error) {

	// set default sender if not
	for i := range msg {
		if msg[i].From == "" {
			msg[i].From = service.From
		}
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return nil, fmt.Errorf("could not encode batch email messages: %w", err)
	}

	// send the data through postmark
	return service.sendMessageThroughPostmark(&buf, "email/batch")
}
