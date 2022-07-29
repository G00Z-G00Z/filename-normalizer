package caseformaters

import "testing"

type specificFormat struct {
	Str    string
	Format CaseFormat
}

var samples = []specificFormat{
	{
		Str:    "hello how are you",
		Format: SpacesCase,
	},
	{
		Str:    "hello_how_are_you",
		Format: SnakeCase,
	},
	{
		Str:    "helloHowAreYou",
		Format: CammelCase,
	},
}

func testWithSpecificFormat(t *testing.T, f *specificFormat, correctTransformation *string, testedFormatData *caseMetaData) {

	result := testedFormatData.Case2snakeCase(f.Str, f.Format)

	if result != *correctTransformation {
		t.Errorf("%s -> %s when using tranforms to %s. Expected %s", f.Str, result, testedFormatData.Name, *correctTransformation)
	}

}

func TestIfEveryFormatIsFound(t *testing.T) {

	for _, sample := range samples {

		_, ok := caseFormattersMetaData[sample.Format]

		if !ok {
			t.Errorf("Format with code %d is not found in metadata!", sample.Format)
		}

	}
}

func TestToSnakeCase(t *testing.T) {

	for _, sampleToBeTested := range samples {

		format := caseFormattersMetaData[sampleToBeTested.Format]

		for _, sampleToTestAgainst := range samples {

			testWithSpecificFormat(t, &sampleToTestAgainst, &sampleToBeTested.Str, &format)

		}

	}

}
