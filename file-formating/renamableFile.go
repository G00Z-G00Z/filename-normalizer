package fileformating

import (
	"fmt"
	"path"
	"strings"

	"github.com/G00Z-G00Z/filename-normalizer/utils"
)

type RenamableFile struct {
	// Name of the file without extension
	Name string
	// Extension of the file with .***
	Ext string

	// Full name of the file
	OriginalName string
}

// Creates a file Renamable File form a string
func CreateRenamableFile(filenameWithExtension string) (RenamableFile, error) {

	ext := path.Ext(filenameWithExtension)
	name := strings.TrimSuffix(filenameWithExtension, ext)

	if len(ext) == 0 && len(name) == 0 {
		dummyFile := RenamableFile{}
		return dummyFile, (fmt.Errorf("tried to format filename: %s but failed", filenameWithExtension))
	}

	file := RenamableFile{
		Name:         name,
		Ext:          ext,
		OriginalName: filenameWithExtension,
	}
	return file, nil
}

func (f *RenamableFile) GetCurrentFormat() utils.CaseFormat {
	return utils.IdentifyCaseFormat(f.Name)
}

func (f *RenamableFile) GetOriginalName() string {
	return f.OriginalName
}

func (f *RenamableFile) GetFullName() string {
	return fmt.Sprintf("%s.%s", f.Name, f.Ext)
}

func (f *RenamableFile) ReturnToOriginal() IRenamable {
	f.Name = f.OriginalName
	return f
}

func (f *RenamableFile) ToCammelCase() IRenamable {
	f.Name = utils.SnakeCase2CammelCase(utils.Spaces2SnakeCase(f.Name))
	return f
}

func (f *RenamableFile) ToSnakeCase() IRenamable {
	f.Name = utils.CammelCase2SnakeCase(utils.Spaces2SnakeCase(f.Name))
	return f
}

func (f *RenamableFile) ToSpaces() IRenamable {
	f.Name = utils.SnakeCase2Spaces(utils.CammelCase2SnakeCase(f.Name))
	return f
}
