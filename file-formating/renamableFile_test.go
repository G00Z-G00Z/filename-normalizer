package fileformating

import (
	"testing"

	"github.com/G00Z-G00Z/filename-normalizer/utils"
)

var (
	cammelCase  = utils.TestUnit[string, string]{Input: "breakingBadFinalDraft.txt"}
	oneWordFile = utils.TestUnit[string, string]{Input: "draft.cpp"}
	snakeCase   = utils.TestUnit[string, string]{Input: "breaking_bad_final_draft.txt"}
	spacesCase  = utils.TestUnit[string, string]{Input: "Breaking bad final draft.txt"}
)

func TestRenamableFileCreation(t *testing.T) {

	filename := "hello"
	extension := ".txt"
	completeName := filename + extension

	renamableFile, err := CreateRenamableFile(completeName)

	if err != nil {
		t.Errorf(err.Error())
	}

	// Test properties
	if filename != renamableFile.Name {
		t.Errorf("filename '%s' doesnt match with file.Name='%s'", filename, renamableFile.Name)
	}

	if extension != renamableFile.Ext {
		t.Errorf("extension '%s' doesnt match with file.Ext='%s'", extension, renamableFile.Ext)
	}

	if completeName != renamableFile.originalName {
		t.Errorf("completeName '%s' doesnt match with file.OriginalName='%s'", completeName, renamableFile.originalName)
	}

	completeName = ""

	renamableFile, err = CreateRenamableFile(completeName)

	if err == nil {
		t.Error("Empty file didnt generate an error")
	}

}

// todo test formatter
func TestReturnToOriginal(t *testing.T) {

	file, _ := CreateRenamableFile(spacesCase.Input)

	spacesCase.Output = file.GetOriginalName()
	spacesCase.Expected = spacesCase.Input

	if !spacesCase.IsCorrect() {
		spacesCase.DisplayError(t)
	}

}

// snakeCaseToCosa = SnakeCase2CammelCase(Space2SnakeCase())
