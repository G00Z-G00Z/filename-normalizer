package caseformaters

import "strings"

// Identifies the case format based on the data map
func IdentifyCaseFormat(s string) CaseFormat {

	for format, data := range caseFormattersMetaData {

		if data.IsCase(s) {
			return format
		}
	}

	return Unkwon
}

//  Identifiers
var isSpaces CaseIdentifier = func(s string) bool {
	return strings.Contains(s, " ")
}

var isSnakeCase CaseIdentifier = func(s string) bool {
	return !isSpaces(s) && strings.Contains(s, "_")
}

var isCammelCase CaseIdentifier = func(s string) bool {
	return !isSpaces(s) && !isSnakeCase(s)
}
