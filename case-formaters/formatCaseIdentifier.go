package caseformaters

import "strings"

func IdentifyCaseFormat(s string) CaseFormat {

	// Has spaces
	if strings.Contains(s, " ") {
		return SpacesCase
	}

	// Has underscores with no spaces
	if strings.Contains(s, "_") {
		return SnakeCase
	}

	return CammelCase
}

//  Identifiers
var isSpaces CaseIdentifier = func(s string) bool {
	return strings.Contains(s, " ")
}

var isSnakeCase CaseIdentifier = func(s string) bool {
	return !isSpaces(s) && strings.Contains(s, "_")
}

var isCammelCase CaseIdentifier = func(s string) bool {
	return !isSpaces(s) && strings.Contains(s, "_")
}
