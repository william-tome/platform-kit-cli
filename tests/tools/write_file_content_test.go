package tools

import (
	"os"
	"testing"

	"github.com/william-tome/platform-kit-cli/tools"
)

func TestWriteFileContent(t *testing.T) {
	testCases := []struct {
		filePath  string
		content   string
		from      string
		expectErr bool
	}{
		{"../testdata/file.txt", "from test file", "TestWriteFileContent", false},
	}

	for _, tc := range testCases {
		err := tools.WriteFileContent(tc.filePath, tc.content, tc.from)
		if tc.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			fileContent, _ := os.ReadFile(tc.filePath)
			if string(fileContent) != tc.content {
				t.Errorf("Expected: %s, but got: %s", tc.content, string(fileContent))
			}
		}
	}
}
