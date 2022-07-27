package utils

import (
	"regexp"
	"strings"
)

var snakeCase2CammelCaseRegex = regexp.MustCompile(`(?m)_(\w)`)

// Transforms spaced out names to snake case
func changeSpacesFor_(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

func Spaces2SnakeCase(s string) string {
	return strings.ToLower(changeSpacesFor_(s))
}

func SnakeCase2CammelCase(s string) string {
	return snakeCase2CammelCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.ToUpper(strings.Trim(m, "_"))
	})
}

func SnakeCase2Spaces(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func CammelCase2SnakeCase(s string) string {
	return SnakeCase2CammelCase(Spaces2SnakeCase(s))
}
