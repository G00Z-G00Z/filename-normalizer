package caseformaters

import (
	"regexp"
	"strings"
)

// Case 2 Snakecase string manipulators
var snake2caseSM = map[CaseFormat]StringFormatter{
	SnakeCase:  snakeCase2SnakeCase,
	CammelCase: snakeCase2CammelCase,
	SpacesCase: snakeCase2Spaces,
}

// Important precompiled regex
var (
	snakeCase2CammelCaseRegex = regexp.MustCompile(`(?m)_(\w)`)
	cammelCase2snakeCaseRegex = regexp.MustCompile(`(?m)[A-Z]`)
)

// String Manipulatos to other cases
func snakeCase2CammelCase(s string) string {
	return snakeCase2CammelCaseRegex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.ToUpper(strings.Trim(m, "_"))
	})
}

func snakeCase2Spaces(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func snakeCase2SnakeCase(s string) string {
	return s
}
