package channel

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFireconnectHeaderOverrides(t *testing.T) {
	cases := []struct {
		name   string
		header http.Header
		want   map[string]string
	}{
		{
			name:   "no fireconnect headers",
			header: http.Header{"Authorization": {"Bearer sk-secret"}, "X-Title": {"Cursor"}},
			want:   map[string]string{},
		},
		{
			name: "harness and version are allow-listed and lower-cased",
			header: http.Header{
				"X-Fireconnect-Harness": {"cursor"},
				"X-Fireconnect-Version": {"1.4.2"},
				"Authorization":         {"Bearer sk-secret"},
			},
			want: map[string]string{
				"x-fireconnect-harness": "cursor",
				"x-fireconnect-version": "1.4.2",
			},
		},
		{
			name: "surrounding whitespace is trimmed",
			header: http.Header{
				"X-Fireconnect-Harness": {"  deepseek  "},
			},
			want: map[string]string{"x-fireconnect-harness": "deepseek"},
		},
		{
			name: "empty values are dropped",
			header: http.Header{
				"X-Fireconnect-Harness": {"claude"},
				"X-Fireconnect-Version": {"   "},
			},
			want: map[string]string{"x-fireconnect-harness": "claude"},
		},
		{
			name:   "nil header",
			header: nil,
			want:   map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, fireconnectHeaderOverrides(c.header))
		})
	}
}

func TestFireconnectHeadersSurviveWithoutPassthroughRules(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://gateway.example/v1/chat/completions", nil)
	assert.NoError(t, err)
	req.Header.Set("X-FireConnect-Harness", "cursor")
	req.Header.Set("X-FireConnect-Version", "1.4.2")

	overrides := fireconnectHeaderOverrides(req.Header)
	applyHeaderOverrideToRequest(req, overrides)

	assert.Equal(t, "cursor", req.Header.Get("x-fireconnect-harness"))
	assert.Equal(t, "1.4.2", req.Header.Get("x-fireconnect-version"))
}
