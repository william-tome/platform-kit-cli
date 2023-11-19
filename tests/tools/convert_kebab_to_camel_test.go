package tools

import (
	"testing"

	"github.com/william.tome/pk-cli/tools"
)

func TestConvertKebabToCamel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello-world", "helloWorld"},
		{"lorem-ipsum-dolor", "loremIpsumDolor"},
	}

	for _, tc := range testCases {
		result := tools.ConvertKebabToCamel(tc.input)
		if result != tc.expected {
			t.Errorf("Expected: %s, but got: %s", tc.expected, result)
		}
	}
}
