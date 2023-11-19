package tools

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ConvertKebabToCamel(str string) string {
	words := strings.Split(str, "-")
	for i := 1; i < len(words); i++ {
		words[i] = cases.Title(language.English).String(words[i])
	}
	return strings.Join(words, "")
}
