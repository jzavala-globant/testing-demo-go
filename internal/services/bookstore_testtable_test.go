package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCoverURL(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedOutput string
	}{
		{
			name:           "Case 1 When empty input, default url output",
			url:            "",
			expectedOutput: defaultCoverURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedOutput, formatCoverURL(tt.url))
		})
	}
}
