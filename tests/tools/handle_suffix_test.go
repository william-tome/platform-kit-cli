package tools_test

import (
	"testing"

	"github.com/william-tome/platform-kit-cli/tools"
)

func TestHandleStringSuffix(t *testing.T) {
	withSuffix := "example.txt"
	withoutSuffix := "example"

	testCases := []struct {
		input          string
		suffixToHandle string
		expected       string
	}{
		{withSuffix, ".txt", withSuffix},
		{withoutSuffix, ".txt", withSuffix},
	}

	for _, tc := range testCases {
		result := tools.HandleStringSuffix(tc.input, tc.suffixToHandle)
		if result != tc.expected {
			t.Errorf("Expected: %s, but got: %s", tc.expected, result)
		}
	}
}
