package route

import (
	"fmt"
	"strings"

	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func appendImportToRouteIndexFile(d *config.RouteProps) error {
	routeName := tools.HandleStringSuffix(d.RouteName, "-routes")
	methodName := tools.ConvertKebabToCamel(routeName)
	fromMethod := "appendImportToRouteIndexFile"

	importInfo := fmt.Sprintf(`import { %s } from '@routes/%s'`, methodName, routeName)

	content, err := tools.ReadFileContent(d.IndexFilePath, fromMethod)
	if err != nil {
		return err
	}

	lastImportPos := strings.LastIndex(string(content), "import {")
	if lastImportPos == -1 {
		return fmt.Errorf("route index file: no import statements found.")
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

func appendExportToRouteIndexFile(d *config.RouteProps) error {
	routeName := tools.HandleStringSuffix(d.RouteName, "-routes")
	methodName := tools.ConvertKebabToCamel(routeName)
	fromMethod := "appendExportToRouteIndexFile"

	content, err := tools.ReadFileContent(d.IndexFilePath, fromMethod)
	if err != nil {
		return err
	}

	startIndex := strings.LastIndex(content, "routes: [")
	endIndex := strings.LastIndex(content, "]")
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		return fmt.Errorf("route index file: no routes found.")
	}

	lastString := content[startIndex:endIndex]
	lastString = strings.TrimSpace(lastString)

	newExport := fmt.Sprintf("%s,\n\t\t%s", lastString, methodName)
	newContent := strings.Replace(content, lastString, newExport, 1)

	if err := tools.WriteFileContent(d.IndexFilePath, newContent, fromMethod); err != nil {
		return err
	}

	return nil
}

func appendToRouteIndex(d *config.RouteProps) error {
	if err := appendImportToRouteIndexFile(d); err != nil {
		return err
	}

	if err := appendExportToRouteIndexFile(d); err != nil {
		return err
	}

	return nil
}
