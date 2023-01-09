package bitbucketapi

import "fmt"

// ToMap convert headers to map[string]string
func (h *Headers) ToMap() map[string]string {
	m := make(map[string]string)
	if h.Token != "" {
		m["Authorization"] = fmt.Sprintf("Bearer %s", h.Token)
	}
	return m
}
