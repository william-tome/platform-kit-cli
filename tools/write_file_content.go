package tools

import (
	"fmt"
	"os"
)

func WriteFileContent(filePath, content, from string) error {
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("write file content - %s: %s", from, err.Error())
	}

	return nil
}
