package fileformating

import (
	"fmt"
	"testing"

	"github.com/G00Z-G00Z/filename-normalizer/utils"
)

var tests []utils.TestUnit[string, string]

func init() {
	tests := []utils.TestUnit[string, string]{}
	fmt.Println(tests)
}

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

	if completeName != renamableFile.OriginalName {
		t.Errorf("completeName '%s' doesnt match with file.OriginalName='%s'", completeName, renamableFile.OriginalName)
	}

}
