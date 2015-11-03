package postmark

import (
	"encoding/base64"
	"encoding/json"
)

// MarshalJSON will encode attachments to json using custom functionality
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
