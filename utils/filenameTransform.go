package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var spaces2snakeCaseRegex = regexp.MustCompile(`(?m) (\w)`)

// Transforms spaced out names to snake case
func spaces2snakeCase(name string) string {
	return spaces2snakeCaseRegex.ReplaceAllStringFunc(name, func(m string) string {
		return fmt.Sprintf("_%s", (strings.Trim(m, " ")))
	})
}
