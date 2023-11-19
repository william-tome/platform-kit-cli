package tools

import (
	"fmt"
	"os"
)

func ReadFileContent(filePath, from string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("read file content - %s: %s", from, err.Error())
	}

	return string(content), nil
}
