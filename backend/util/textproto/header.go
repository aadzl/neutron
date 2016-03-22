package textproto

import (
	"net/textproto"
	"time"

	"github.com/emersion/neutron/backend"
)

func GetMessageHeader(msg *backend.Message) textproto.MIMEHeader {
	h := textproto.MIMEHeader{}

	h.Set("MIME-Version", "1")

	h.Set("Subject", msg.Subject)
	h.Set("From", FormatEmail(msg.Sender))
	h.Set("Date", time.Unix(msg.Time, 0).Format(time.RFC1123Z))

	for _, to := range msg.ToList {
		h.Add("To", FormatEmail(to))
	}
	for _, cc := range msg.CCList {
		h.Add("Cc", FormatEmail(cc))
	}

	if msg.ReplyTo != nil {
		h.Set("Reply-To", FormatEmail(msg.ReplyTo))
	}

	return h
}

func GetOutgoingMessageHeader(msg *backend.OutgoingMessage) textproto.MIMEHeader {
	h := GetMessageHeader(msg.Message)

	if msg.InReplyTo != "" {
		h.Set("In-Reply-To", msg.InReplyTo)
	}
	if msg.References != "" {
		h.Set("References", msg.References)
	}

	return h
}

func FomatHeader(h textproto.MIMEHeader) string {
	output := ""
	for key, values := range h {
		for _, value := range values {
			output += key + ": " + value + "\r\n"
		}
	}
	return output
}

func formatMessage(header textproto.MIMEHeader, body string) string {
	return FomatHeader(header) + "\r\n" + body
}

func FormatMessage(msg *backend.Message) string {
	header := GetMessageHeader(msg)
	return formatMessage(header, msg.Body)
}

func FormatOutgoingMessage(msg *backend.OutgoingMessage) string {
	header := GetOutgoingMessageHeader(msg)
	return formatMessage(header, msg.MessagePackage.Body)
}