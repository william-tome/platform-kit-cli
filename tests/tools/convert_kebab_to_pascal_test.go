package tools

import (
	"testing"

	"github.com/william.tome/pk-cli/tools"
)

func TestConvertKebabToPascal(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello-world", "HelloWorld"},
		{"lorem-ipsum-dolor", "LoremIpsumDolor"},
	}

	for _, tc := range testCases {
		result := tools.ConvertKebabToPascal(tc.input)
		if result != tc.expected {
			t.Errorf("Expected: %s, but got: %s", tc.expected, result)
		}
	}
}
