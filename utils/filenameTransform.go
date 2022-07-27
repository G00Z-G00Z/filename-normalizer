package utils

import (
	"regexp"
	"strings"
)

var changeSpacesFor_Regex = regexp.MustCompile(`(?m) `)
var snakeCase2CammelCaseRegex = regexp.MustCompile(`(?m)_(\w)`)

const subSpaces = "_"

// Transforms spaced out names to snake case
func changeSpacesFor_(s string) string {
	return changeSpacesFor_Regex.ReplaceAllString(s, subSpaces)
}

func Spaces2SnakeCase(s string) string {
	return strings.ToLower(changeSpacesFor_(s))
}

func SnakeCase2CammelCase(s string) string {
	return snakeCase2CammelCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.ToUpper(strings.Trim(m, "_"))
	})

}
