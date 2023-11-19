package tools

import "strings"

func HandleStringSuffix(str, suffixToHandle string) string {
	if strings.HasSuffix(str, suffixToHandle) {
		return str
	}

	return str + suffixToHandle
}
