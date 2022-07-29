package caseformaters

import (
	"fmt"
	"strings"
)

// Case 2 Snakecase string manipulators
var case2snakeCaseSM = map[CaseFormat]StringFormatter{
	SnakeCase:  snakeCase2SnakeCase,
	CammelCase: cammelCase2SnakeCase,
	SpacesCase: spaces2SnakeCase,
}

// String Manipulators to Snake Case
func spaces2SnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "_"))
}

func cammelCase2SnakeCase(s string) string {
	return cammelCase2snakeCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return fmt.Sprintf("_%s", strings.ToLower(m))
	})
}
