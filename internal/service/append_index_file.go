package service

import (
	"fmt"
	"strings"

	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func appendImportToServiceIndexFile(d *config.ServiceProps) error {
	methodName := tools.ConvertKebabToCamel(d.ServiceName)
	importInfo := fmt.Sprintf(`import { %s } from './%s'`, methodName, d.ServiceName)
	fromMethod := "appendImportToServiceIndexFile"

	content, err := tools.ReadFileContent(d.IndexFilePath, fromMethod)
	if err != nil {
		return err
	}

	lastImportPos := strings.LastIndex(string(content), "import")
	if lastImportPos == -1 {
		return fmt.Errorf("service index file: no import statements found.")
	}

	endOfLinePos := strings.IndexByte(string(content[lastImportPos:]), '\n')
	if endOfLinePos == -1 {
		endOfLinePos = len(content)
	} else {
		endOfLinePos += lastImportPos
	}

	prevContent := string(content[:endOfLinePos])
	nextContent := string(content[endOfLinePos:])

	newContent := fmt.Sprintf("%s\n%s%s", prevContent, importInfo, nextContent)

	if err := tools.WriteFileContent(d.IndexFilePath, newContent, fromMethod); err != nil {
		return err
	}

	return nil
}

func appendExportToServiceIndexFile(d *config.ServiceProps) error {
	methodName := tools.ConvertKebabToCamel(d.ServiceName)
	fromMethod := "appendExportToServiceIndexFile"

	content, err := tools.ReadFileContent(d.IndexFilePath, fromMethod)
	if err != nil {
		return err
	}

	startIndex := strings.LastIndex(content, "export {")
	endIndex := strings.LastIndex(content, "}")
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		return fmt.Errorf("service index file: no export statements found.")
	}

	lastString := content[startIndex:endIndex]
	lastString = strings.TrimSpace(lastString)

	newExport := fmt.Sprintf("%s,\n\t%s", lastString, methodName)
	newContent := strings.Replace(content, lastString, newExport, 1)

	if err := tools.WriteFileContent(d.IndexFilePath, newContent, fromMethod); err != nil {
		return err
	}

	return nil
}

func appendServiceToIndex(d *config.ServiceProps) error {
	if err := appendImportToServiceIndexFile(d); err != nil {
		return err
	}

	if err := appendExportToServiceIndexFile(d); err != nil {
		return err
	}

	return nil
}
