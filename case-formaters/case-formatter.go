package caseformaters

import "fmt"

// Case 2 Snakecase string manipulators
var case2snakeCaseSM = map[CaseFormat]StringFormatter{
	SnakeCase:  func(s string) string { return s },
	CammelCase: cammelCase2SnakeCase,
	SpacesCase: spaces2SnakeCase,
}

// Normalizes any case to snake case, for later processing
// Uses a map to store the functions used to transform any case to SnakeCase
var normalizeToSnakeCase CaseFormatter = func(s string, currentFormat CaseFormat) string {

	fn, ok := case2snakeCaseSM[currentFormat]

	if !ok {
		panic(fmt.Sprintf("an unrecognizable case format with code '%d' doesnt have a CaseToSnakeCase function asigned!!", currentFormat))
	}

	return fn(s)

}

// Decorator for checking case and optimizing transforms
// It requires a function that transforms snakeCase to any other format
// It also normalizes the input to snakeCase if needed
var stringFormatter2CaseFormatter = func(targetCase CaseFormat) CaseFormatter {
	return func(s string, currentFormat CaseFormat) string {
		if currentFormat == targetCase {
			return s
		}
		return snake2caseSM[targetCase](normalizeToSnakeCase(s, currentFormat))
	}

}

// Case Transforms

// Transforms strings to snakeCase
var ToSnakeCase CaseFormatter = stringFormatter2CaseFormatter(SnakeCase)

// Transforms strings to cammelCase
var ToCammelCase CaseFormatter = stringFormatter2CaseFormatter(CammelCase)

// Transforms strings to spacesCase
var ToSpaces CaseFormatter = stringFormatter2CaseFormatter(SpacesCase)
