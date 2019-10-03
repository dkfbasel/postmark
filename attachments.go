package postmark

import (
	"encoding/base64"
	"encoding/json"
)

// Attachment should be used to attach files to the emails. Note that the content
// will be base64 encoded to be sent to the client
type Attachment struct {
	Name        string
	ContentType string
	Content     []byte
}

// MarshalJSON will encode attachments to json using custom encoding functionality
func (attachment *Attachment) MarshalJSON() ([]byte, error) {

	preparedForMarshalling := &struct {
		Name        string
		Content     string
		ContentType string
	}{
		Name:        attachment.Name,
		ContentType: attachment.ContentType,
	}

	// use base64 encoding to send mail attachments
	preparedForMarshalling.Content = base64.StdEncoding.EncodeToString(attachment.Content)

	return json.Marshal(preparedForMarshalling)
}
