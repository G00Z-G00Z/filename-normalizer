package utils

import (
	"regexp"
	"strings"
)

var changeSpacesFor_Regex = regexp.MustCompile(`(?m) `)

const substitution = "_"

// Transforms spaced out names to snake case
func changeSpacesFor_(name string) string {
	return changeSpacesFor_Regex.ReplaceAllString(name, substitution)
}

func Spaces2SnakeCase(s string) string {
	return strings.ToLower(changeSpacesFor_(s))
}
