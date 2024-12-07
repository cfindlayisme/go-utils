package strings_test

import (
	"testing"

	"github.com/cfindlayisme/go-utils/strings"
)

func TestStripNewlines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No newlines",
			input:    "Hello, World!",
			expected: "Hello, World!",
		},
		{
			name:     "Newlines only",
			input:    "\n\n\r\r",
			expected: "",
		},
		{
			name:     "Mixed content with newlines",
			input:    "Hello\nWorld\r!",
			expected: "HelloWorld!",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Newlines at start and end",
			input:    "\nHello, World!\r",
			expected: "Hello, World!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := strings.StripNewlines(test.input)
			if actual != test.expected {
				t.Errorf("StripNewlines(%q) = %q; want %q", test.input, actual, test.expected)
			}
		})
	}
}
