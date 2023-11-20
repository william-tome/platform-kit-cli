package tools_test

import (
	"os"
	"testing"

	"github.com/william-tome/platform-kit-cli/tools"
)

func TestReadFileContent(t *testing.T) {
	testCases := []struct {
		filePath  string
		from      string
		expected  string
		expectErr bool
	}{
		{"../testdata/file.txt", "TestReadFileContent", "expected file content", false},
	}

	for _, tc := range testCases {
		content, err := tools.ReadFileContent(tc.filePath, tc.from)
		if tc.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			fileContent, _ := os.ReadFile(tc.filePath)
			if content != string(fileContent) {
				t.Errorf("Expected: %s, but got: %s", string(fileContent), content)
			}
		}
	}
}
