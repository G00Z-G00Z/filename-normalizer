package caseformaters

import (
	"testing"
)

func TestSnakeCaseToSpaces(t *testing.T) {
	for _, test := range testSnakeCase2Spaces {
		test.SetOutput(snakeCase2Spaces(test.Input))
		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestSnakeCase2CammelCase(t *testing.T) {
	for _, test := range testSnakeCase2CammelCase {
		test.SetOutput(snakeCase2CammelCase(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}
