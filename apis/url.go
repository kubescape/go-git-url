package apis

import (
	"net/url"
	"strings"
)

// HostWithPort returns the URL host, stripping the port if it is the
// default for the scheme (so detection-friendly hostnames are used for
// canonical URLs) while preserving non-default ports for self-hosted
// instances where the port is load-bearing for API calls.
func HostWithPort(parsedURL *url.URL) string {
	if parsedURL == nil {
		return ""
	}

	port := parsedURL.Port()
	if port == "" || isDefaultPort(parsedURL.Scheme, port) {
		return preserveIPv6Brackets(parsedURL.Hostname())
	}
	return parsedURL.Host
}

func preserveIPv6Brackets(host string) string {
	if strings.ContainsRune(host, ':') {
		return "[" + host + "]"
	}
	return host
}

func isDefaultPort(scheme, port string) bool {
	switch scheme {
	case "https", "wss":
		return port == "443"
	case "http", "ws":
		return port == "80"
	case "ssh", "git+ssh":
		return port == "22"
	case "git":
		return port == "9418"
	}
	return false
}
