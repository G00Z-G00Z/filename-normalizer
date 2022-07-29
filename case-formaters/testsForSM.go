package caseformaters

import "github.com/G00Z-G00Z/filename-normalizer/utils"

var (
	testToSnakeCase []utils.TestUnit[string, string] = []utils.TestUnit[string, string]{
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
	testSnakeCase2CammelCase []utils.TestUnit[string, string] = []utils.TestUnit[string, string]{
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
	testSnakeCase2Spaces []utils.TestUnit[string, string] = []utils.TestUnit[string, string]{
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

	testCammelCase2SnakeCase []utils.TestUnit[string, string] = []utils.TestUnit[string, string]{
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
)
