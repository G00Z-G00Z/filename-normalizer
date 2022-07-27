package utils

import (
	"regexp"
)

var changeSpacesFor_Regex = regexp.MustCompile(`(?m) `)

const substitution = "_"

// Transforms spaced out names to snake case
func changeSpacesFor_(name string) string {
	return changeSpacesFor_Regex.ReplaceAllString(name, substitution)
}
