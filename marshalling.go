package postmark

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

// MarshalJSON will encode a message to json using custom functionality
func (message *Message) MarshalJSON() ([]byte, error) {

	preparedForMarshalling := &struct {
		From string `json:",omitempty"`
		To   string `json:",omitempty"`

		Cc  string `json:",omitempty"`
		Bcc string `json:",omitempty"`

		Subject string `json:",omitempty"`
		Tag     string `json:",omitempty"`

		HTMLBody string `json:"HtmlBody,omitempty"`
		TextBody string `json:",omitempty"`

		ReplyTo string `json:",omitempty"`

		Headers     []map[string]string `json:",omitempty"`
		Attachments []Attachment        `json:",omitempty"`
	}{
		From: message.From,
		To:   addressesToList(message.To),
		Cc:   addressedToList(message.Cc),
		Bcc:  addressedToList(message.Cc),

		Subject: message.Subject,
		Tag:     message.Tag,

		HTMLBody: message.HTMLBody,
		TextBody: message.TextBody,

		ReplyTo: addressedToList(message.Cc),

		Headers:     transformMailHeaders(message.Headers),
		Attachments: message.Attachments,
	}

	return json.Marshal(preparedForMarshalling)
}

// // MarshalJSON will encode messageto json using custom functionality
// func (message *MessageWithTemplate) MarshalJSON() ([]byte, error) {
// 	preparedForMarshalling := &struct {
// 		From string `json:",omitempty"`
// 		To   string `json:",omitempty"`
//
// 		Cc  string `json:",omitempty"`
// 		Bcc string `json:",omitempty"`
//
// 		Subject string `json:",omitempty"`
// 		Tag     string `json:",omitempty"`
//
// 		TemplateID    int                    `json:"TemplateId,omitempty"`
// 		TemplateModel map[string]interface{} `json:",omitempty"`
//
// 		ReplyTo string `json:",omitempty"`
//
// 		Headers     []map[string]string `json:",omitempty"`
// 		Attachments []Attachment        `json:",omitempty"`
// 	}{
// 		From: message.From,
// 		To:   addressesToList(message.To),
// 		Cc:   addressedToList(message.Cc),
// 		Bcc:  addressedToList(message.Cc),
//
// 		Subject: message.Subject,
// 		Tag:     message.Tag,
//
// 		TemplateID:    message.TemplateID,
// 		TemplateModel: message.TemplateModel,
//
// 		ReplyTo: addressedToList(message.Cc),
//
// 		Headers:     transformMailHeaders(message.Headers),
// 		Attachments: message.Attachments,
// 	}
//
// 	return json.Marshal(preparedForMarshalling)
// }

// MarshalJSON will encode attachments to json using custom functionality
func (attachment *Attachment) MarshalJSON() ([]byte, error) {

	preparedForMarshalling := &struct {
		Name        string
		Content     string
		ContentType string
	}{}

	preparedForMarshalling.Name = attachment.Name

	// read the attachment content
	content, err := ioutil.ReadAll(attachment.Content)
	if err != nil {
		return nil, err
	}

	preparedForMarshalling.Content = base64.StdEncoding.EncodeToString(content)
	preparedForMarshalling.ContentType = attachment.ContentType

	return json.Marshal(doc)
}
