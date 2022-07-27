package utils

import (
	"testing"
)

var (
	testForRemovingSpaces    []TestUnit[string, string]
	testToSnakeCase          []TestUnit[string, string]
	testSnakeCase2CammelCase []TestUnit[string, string]
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

}

func TestChangeSpacesFor_(t *testing.T) {
	for _, test := range testForRemovingSpaces {
		test.SetOutput(changeSpacesFor_(test.Input))

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
