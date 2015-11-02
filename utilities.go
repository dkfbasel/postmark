package postmark

import (
	"net/mail"
	"net/url"
	"strings"
)

// addressesToList will join a list of addresses to a comma separated list
func addressesToList(addresses []*mail.Address) string {

	list := []string{}

	for _, address := range addresses {
		list = append(list, address.String())
	}

	return strings.Join(list, ", ")
}

func transformMailHeaders(mailHeaders mail.Header) []map[string]string {

	headers := []map[string]string{}

	for k, vs := range mailHeaders {
		for _, v := range vs {
			headers = append(headers, map[string]string{
				"Name":  k,
				"Value": v,
			})
		}
	}

	return headers
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
