package utils

import (
	"testing"
)

var (
	testForRemovingSpaces    []TestUnit[string, string]
	testToSnakeCase          []TestUnit[string, string]
	testSnakeCase2CammelCase []TestUnit[string, string]
	testSnakeCase2Spaces     []TestUnit[string, string]
	testCammelCase2SnakeCase []TestUnit[string, string]
)

func init() {

	testForRemovingSpaces = []TestUnit[string, string]{
		{
			Input:    "hey how are you",
			Expected: "hey_how_are_you",
		}, {
			Input:    "I Dont LIke Files wiTh spaces",
			Expected: "I_Dont_LIke_Files_wiTh_spaces",
		}, {
			Input:    "heLLoYou",
			Expected: "heLLoYou",
		},
		{
			Input:    "hello",
			Expected: "hello",
		},
	}

	testToSnakeCase = []TestUnit[string, string]{
		{
			Input:    "hey how are you",
			Expected: "hey_how_are_you",
		}, {
			Input:    "I Dont LIke Files wiTh spaces",
			Expected: "i_dont_like_files_with_spaces",
		}, {
			Input:    "heLLoYou",
			Expected: "helloyou",
		},
		{
			Input:    "hello",
			Expected: "hello",
		},
	}

	testSnakeCase2CammelCase = []TestUnit[string, string]{
		{
			Input:    "hello",
			Expected: "hello",
		}, {
			Input:    "i_am_not_a_king",
			Expected: "iAmNotAKing",
		}, {
			Input:    "just_do_it",
			Expected: "justDoIt",
		}, {
			Input:    "hello_how_are_you",
			Expected: "helloHowAreYou",
		},
	}

	testSnakeCase2Spaces = []TestUnit[string, string]{
		{
			Expected: "hey how are you",
			Input:    "hey_how_are_you",
		},
		{
			Expected: "i dont like files with spaces",
			Input:    "i_dont_like_files_with_spaces",
		},
		{
			Expected: "helloyou",
			Input:    "helloyou",
		},
		{
			Expected: "hello",
			Input:    "hello",
		},
	}

	testCammelCase2SnakeCase = []TestUnit[string, string]{
		{
			Expected: "hello",
			Input:    "hello",
		}, {
			Expected: "i_am_not_a_king",
			Input:    "iAmNotAKing",
		}, {
			Expected: "just_do_it",
			Input:    "justDoIt",
		}, {
			Expected: "hello_how_are_you",
			Input:    "helloHowAreYou",
		},
	}

}

func TestChangeSpacesFor_(t *testing.T) {
	for _, test := range testForRemovingSpaces {
		test.SetOutput(changeSpacesFor_(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestSnakeCaseToSpaces(t *testing.T) {
	for _, test := range testSnakeCase2Spaces {
		test.SetOutput(SnakeCase2Spaces(test.Input))
		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestSpaces2SnakeCase(t *testing.T) {
	for _, test := range testToSnakeCase {
		test.SetOutput(Spaces2SnakeCase(test.Input))
		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestSnakeCase2CammelCase(t *testing.T) {
	for _, test := range testSnakeCase2CammelCase {
		test.SetOutput(SnakeCase2CammelCase(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}

func TestCammelCase2SnakeCase(t *testing.T) {
	for _, test := range testCammelCase2SnakeCase {
		test.SetOutput(CammelCase2SnakeCase(test.Input))

		if !test.IsCorrect() {
			test.DisplayError(t)
		}
	}
}
