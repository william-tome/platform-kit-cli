package tools

import "strings"

func HandleImportPrefix(str, prefixToHandle string) string {
	if strings.HasPrefix(str, prefixToHandle) {
		return str
	}

	return strings.Replace(str, "src", prefixToHandle, 1)
}
