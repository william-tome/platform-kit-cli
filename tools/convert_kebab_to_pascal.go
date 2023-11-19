package tools

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ConvertKebabToPascal(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = cases.Title(language.English).String(str)

	return strings.ReplaceAll(str, " ", "")
}
