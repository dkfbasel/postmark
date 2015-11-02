package postmark

import "net/mail"

const (
	// DefaultHost is used to hold the standard config to send messages to
	DefaultHost = "api.postmarkapp.com"
)

// Message will contain the basic information for an email message
type Message struct {
	From string   `json:",omitempty"`
	To   []string `json:",omitempty"`

	Cc  []string `json:",omitempty"`
	Bcc []string `json:",omitempty"`

	Subject string `json:",omitempty"`
	Tag     string `json:",omitempty"`

	HTMLBody string `json:"HtmlBody,omitempty"`
	TextBody string `json:",omitempty"`

	ReplyTo string `json:",omitempty"`

	Headers     mail.Header  `json:",omitempty"`
	Attachments []Attachment `json:",omitempty"`
}

// MessageWithTemplate can be used to send a message using a template
type MessageWithTemplate struct {
	Message

	TemplateID    int `json:"TemplateID"`
	TemplateModel map[string]interface{}
}

// Attachment should be used to attach files to the emails. Note that the content
// will be base64 encoded to be sent to the client
type Attachment struct {
	Name        string
	ContentType string
	Content     []byte
}
