package caseformaters

import (
	"testing"

	"github.com/G00Z-G00Z/filename-normalizer/utils"
)

type caseTests struct {
	Spaces     string
	SnakeCase  string
	CammelCase string
}

type inputWithCurrentCase struct {
	Input       string
	CurrentCase CaseFormat
}

var testsCases = []utils.TestUnit[caseTests, caseTests]{
	{
		Input: caseTests{
			Spaces:     "hello how are you",
			SnakeCase:  "hello_how_are_you",
			CammelCase: "helloHowAreYou",
		},
	},
}

func testFormatTransform(t *testing.T, str string, fn CaseFormatter, caseFormat CaseFormat, correct string) {

	if transformed := fn(str, caseFormat); transformed != correct {
		t.Errorf("%v fn transformed %s -> %s. It supposed to transfrom into %s\n", fn, str, transformed, correct)
	}

}

func testSeveralFormatTransforms(t *testing.T, correctTransform string, testingFunc CaseFormatter, inputs []inputWithCurrentCase) {
	for _, input := range inputs {
		testFormatTransform(t, input.Input, testingFunc, input.CurrentCase, correctTransform)
	}

}

func TestCaseToSnakeCase(t *testing.T) {

}
