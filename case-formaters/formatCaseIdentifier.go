package caseformaters

import "strings"

func IdentifyCaseFormat(s string) CaseFormat {

	// Has spaces
	if strings.Contains(s, " ") {
		return Spaces
	}

	// Has underscores with no spaces
	if strings.Contains(s, "_") {
		return SnakeCase
	}

	return CammelCase
}
