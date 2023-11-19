package tools_test

import (
	"testing"

	"github.com/william-tome/platform-kit-cli/tools"
)

func TestHandleImportPrefix(t *testing.T) {
	withSrcPrefix := "src/example/test.go"
	withoutSrcPrefix := "@/example/test.go"

	testCases := []struct {
		input          string
		prefixToHandle string
		expected       string
	}{
		{withSrcPrefix, "@", withoutSrcPrefix},
		{withoutSrcPrefix, "@", withoutSrcPrefix},
	}

	for _, tc := range testCases {
		result := tools.HandleImportPrefix(tc.input, tc.prefixToHandle)
		if result != tc.expected {
			t.Errorf("Expected: %s, but got: %s", tc.expected, result)
		}
	}
}
