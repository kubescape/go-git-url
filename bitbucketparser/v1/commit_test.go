package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseRawUser(t *testing.T) {
	tests := []struct {
		raw       string
		wantName  string
		wantEmail string
	}{
		{
			raw:       "David Wertenteil <dwertent@cyberarmor.io>",
			wantName:  "David Wertenteil",
			wantEmail: "dwertent@cyberarmor.io",
		},
		{
			raw:       "github-actions[bot] <github-actions[bot]@users.noreply.github.com>",
			wantName:  "github-actions[bot]",
			wantEmail: "github-actions[bot]@users.noreply.github.com",
		},
		{
			raw:       "David Wertenteil",
			wantName:  "David Wertenteil",
			wantEmail: "",
		},
		{
			raw:       "<dwertent@cyberarmor.io>",
			wantName:  "",
			wantEmail: "dwertent@cyberarmor.io",
		},
	}
	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			got, got1 := parseRawUser(tt.raw)
			assert.Equalf(t, tt.wantName, got, "parseRawUser(%v)", tt.raw)
			assert.Equalf(t, tt.wantEmail, got1, "parseRawUser(%v)", tt.raw)
		})
	}
}
