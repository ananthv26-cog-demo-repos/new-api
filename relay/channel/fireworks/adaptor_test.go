package fireworks

import (
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/relay/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequestURL(t *testing.T) {
	cases := []struct {
		name      string
		baseURL   string
		relayMode int
		expected  string
	}{
		{
			name:      "default base url with version suffix",
			baseURL:   "https://api.fireworks.ai/inference/v1",
			relayMode: constant.RelayModeChatCompletions,
			expected:  "https://api.fireworks.ai/inference/v1/chat/completions",
		},
		{
			name:      "base url without version suffix",
			baseURL:   "https://api.fireworks.ai/inference",
			relayMode: constant.RelayModeChatCompletions,
			expected:  "https://api.fireworks.ai/inference/v1/chat/completions",
		},
		{
			name:      "base url with trailing slash",
			baseURL:   "https://api.fireworks.ai/inference/v1/",
			relayMode: constant.RelayModeChatCompletions,
			expected:  "https://api.fireworks.ai/inference/v1/chat/completions",
		},
		{
			name:      "completions",
			baseURL:   "https://api.fireworks.ai/inference/v1",
			relayMode: constant.RelayModeCompletions,
			expected:  "https://api.fireworks.ai/inference/v1/completions",
		},
		{
			name:      "embeddings",
			baseURL:   "https://api.fireworks.ai/inference/v1",
			relayMode: constant.RelayModeEmbeddings,
			expected:  "https://api.fireworks.ai/inference/v1/embeddings",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			info := &relaycommon.RelayInfo{
				RelayMode:   tc.relayMode,
				ChannelMeta: &relaycommon.ChannelMeta{ChannelBaseUrl: tc.baseURL},
			}

			url, err := (&Adaptor{}).GetRequestURL(info)

			require.NoError(t, err)
			assert.Equal(t, tc.expected, url)
		})
	}
}
