package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Catches all the _W pairs
var snakeCase2CammelCaseRegex = regexp.MustCompile(`(?m)_(\w)`)

// Catches all uppercase letters
var cammelCase2snakeCaseRegex = regexp.MustCompile(`(?m)[A-Z]`)

// Transforms spaced out names to snake case
func changeSpacesFor_(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// Transforms a string with spaces into a snake case format
func Spaces2SnakeCase(s string) string {
	return strings.ToLower(changeSpacesFor_(s))
}

// Transforms snake case format to cammel case format
func SnakeCase2CammelCase(s string) string {
	return snakeCase2CammelCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.ToUpper(strings.Trim(m, "_"))
	})
}

// Transforms snake case format to spaces
func SnakeCase2Spaces(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

// Transforms cammel case format to snake case format
func CammelCase2SnakeCase(s string) string {
	return cammelCase2snakeCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return fmt.Sprintf("_%s", strings.ToLower(m))
	})
}
