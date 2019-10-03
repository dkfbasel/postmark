package postmark

// Message will contain the basic information for an email message
type Message struct {
	From string `json:",omitempty"`
	To   string `json:",omitempty"`

	Cc  string `json:",omitempty"`
	Bcc string `json:",omitempty"`

	Subject string `json:",omitempty"`
	Tag     string `json:",omitempty"`

	HTMLBody string `json:"HtmlBody,omitempty"`
	TextBody string `json:"TextBody,omitempty"`

	ReplyTo string `json:",omitempty"`

	TrackOpens bool   `json:",omitempty"`
	TrackLinks string `json:",omitempty"`

	Headers     []MailHeader `json:",omitempty"`
	Attachments []Attachment `json:",omitempty"`
}

// MessageWithTemplate can be used to send a message using a specific
// postmark template
type MessageWithTemplate struct {
	Message

	TemplateID    int    `json:"TemplateId,omitempty"`
	TemplateAlias string `json:",omitempty"`
	TemplateModel map[string]interface{}
}

// Tracking of links in messages
const TrackNone string = "None"
const TrackHTMLAndText string = "HtmlAndText"
const TrackHtmlOnly string = "HtmlOnly"
const TrackTextOnly string = "TextOnly"

// MailHeader can be used to set specific mail headers
type MailHeader struct {
	Name  string `json:",omitempty"`
	Value string
}
