package model

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientTagFromRequest(t *testing.T) {
	cases := []struct {
		name   string
		header http.Header
		want   string
	}{
		{
			name:   "no attribution headers",
			header: http.Header{"X-Title": {"Cursor"}},
			want:   "",
		},
		{
			name: "harness and version",
			header: http.Header{
				"X-Fireconnect-Harness": {"cursor"},
				"X-Fireconnect-Version": {"1.4.2"},
			},
			want: "cursor/1.4.2",
		},
		{
			name:   "harness without version",
			header: http.Header{"X-Fireconnect-Harness": {"deepseek"}},
			want:   "deepseek",
		},
		{
			name: "unknown version is kept verbatim",
			header: http.Header{
				"X-Fireconnect-Harness": {"claude"},
				"X-Fireconnect-Version": {"unknown"},
			},
			want: "claude/unknown",
		},
		{
			name: "version without harness is not tagged",
			header: http.Header{
				"X-Fireconnect-Version": {"1.4.2"},
			},
			want: "",
		},
		{
			name: "values are trimmed",
			header: http.Header{
				"X-Fireconnect-Harness": {"  codex  "},
				"X-Fireconnect-Version": {"  2.0.0  "},
			},
			want: "codex/2.0.0",
		},
		{
			name: "oversized tag is truncated to the column width",
			header: http.Header{
				"X-Fireconnect-Harness": {strings.Repeat("a", 200)},
			},
			want: strings.Repeat("a", maxClientTagLength),
		},
		{
			name: "truncation never splits a multi-byte character",
			header: http.Header{
				// The 128-byte cut lands mid-rune: 126 ASCII bytes plus a 3-byte rune.
				"X-Fireconnect-Harness": {strings.Repeat("a", 126) + strings.Repeat("界", 10)},
			},
			want: strings.Repeat("a", 126),
		},
		{
			name:   "nil header",
			header: nil,
			want:   "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, ClientTagFromRequest(c.header))
		})
	}
}
