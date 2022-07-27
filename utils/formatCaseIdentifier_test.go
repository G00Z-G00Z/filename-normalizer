package utils

import "testing"

var tests []TestUnit[string, CaseFormat]

func init() {

	tests = []TestUnit[string, CaseFormat]{
		{
			Input:    "hello who are you",
			Expected: Spaces,
		},
		{
			Input:    "my name is_not yours",
			Expected: Spaces,
		},
		{
			Input:    "i_am_in_snake_case",
			Expected: SnakeCase,
		},
		{
			Input:    "amInCammelCase",
			Expected: CammelCase,
		},
		{
			Input:    "yikes",
			Expected: CammelCase,
		},
	}

}

func TestIdentifyCaseFormat(t *testing.T) {

	for _, test := range tests {
		test.SetOutput(IdentifyCaseFormat(test.Input))
		if !test.IsCorrect() {
			test.DisplayError(t)
		}

	}

}
