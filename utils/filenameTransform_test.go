package utils

import "testing"

var tests []TestUnit[string, string]

func init() {

	tests = []TestUnit[string, string]{
		{
			Input:    "hey how are you.txt",
			Expected: "hey_how_are_you.txt",
		}, {
			Input:    "I Dont LIke Files wiTh spaces.txt",
			Expected: "I_Dont_LIke_Files_wiTh_spaces.txt",
		}, {
			Input:    "heLLoYou.EXE",
			Expected: "heLLoYou.EXE",
		},
		{
			Input:    "hello.txt",
			Expected: "hello.txt",
		},
	}

}

func TestSpaces2snakeCase(t *testing.T) {
	for _, test := range tests {
		test.SetOutput(spaces2snakeCase(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}

	}
}
