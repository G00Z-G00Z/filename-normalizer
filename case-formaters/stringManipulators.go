package caseformaters

import (
	"fmt"
	"regexp"
	"strings"
)

// Case 2 Snakecase string manipulators
var snake2caseSM = map[CaseFormat]StringFormatter{
	SnakeCase:  func(s string) string { return s },
	CammelCase: snakeCase2CammelCase,
	SpacesCase: snakeCase2Spaces,
}

// Important precompiled regex
var (
	snakeCase2CammelCaseRegex = regexp.MustCompile(`(?m)_(\w)`)
	cammelCase2snakeCaseRegex = regexp.MustCompile(`(?m)[A-Z]`)
)

// String Manipulators to Snake Case
func spaces2SnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "_"))
}

func cammelCase2SnakeCase(s string) string {
	return cammelCase2snakeCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return fmt.Sprintf("_%s", strings.ToLower(m))
	})
}

// String Manipulatos to other cases
func snakeCase2CammelCase(s string) string {
	return snakeCase2CammelCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.ToUpper(strings.Trim(m, "_"))
	})
}

func snakeCase2Spaces(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}
