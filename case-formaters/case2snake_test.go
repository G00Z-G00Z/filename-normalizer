package caseformaters

import (
	"testing"
)

func TestSpaces2SnakeCase(t *testing.T) {
	for _, test := range testToSnakeCase {
		test.SetOutput(spaces2SnakeCase(test.Input))
		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestCammelCase2SnakeCase(t *testing.T) {
	for _, test := range testCammelCase2SnakeCase {
		test.SetOutput(cammelCase2SnakeCase(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}
