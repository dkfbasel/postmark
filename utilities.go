package postmark

import (
	"net/mail"
	"strings"
)

// Address will return a from/to compatible string of the given contact
func Address(name, email string) string {
	address := mail.Address{
		Name:    name,
		Address: email,
	}
	return address.String()
}

// Addresses will join a list of addresses to a comma separated list
func Addresses(addresses ...mail.Address) string {
	list := []string{}
	for _, address := range addresses {
		list = append(list, address.String())
	}
	return strings.Join(list, ", ")
}

// Emails will join a list of email addresses into a comma separated list
func Emails(emails ...string) string {
	return strings.Join(emails, ", ")
}
