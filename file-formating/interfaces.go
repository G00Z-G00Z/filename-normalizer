package fileformating

import (
	"fmt"
	"strings"

	"github.com/G00Z-G00Z/filename-normalizer/utils"
)

// Interface of a Renamable element
type IRenamable interface {
	// Gets the full name of the file
	GetFullName() string

	// Get the current format
	GetCurrentFormat() utils.CaseFormat

	// Get the unchanged name when first instanciated
	GetOriginalName() string

	// Revert back to original name
	ReturnToOriginal()

	// Formatters
	ChangeCase(formatter IFormatter)
}

// Formatter that takes a string and transform it
type IFormatter interface {
	Transform(s string) string
}

type BaseFormatter struct{}

func (t *BaseFormatter) Transform(s string) string {
	return s
}

type UpperCaseFormatter struct {
	otherFormatter IFormatter
}

type SuffixFormatter struct {
	otherFormatter IFormatter
}

func (t *UpperCaseFormatter) Transform(s string) string {
	s = t.otherFormatter.Transform(s)
	s = strings.ToUpper(s)
	return s
}

func (t *SuffixFormatter) Transform(s string) string {
	s = t.otherFormatter.Transform(s)
	s = "suf." + s
	return s
}

var UpperSuffix IFormatter

func init() {

	fmt.Println("Hola estoy corriendo init")

	UpperSuffix =
		&SuffixFormatter{
			otherFormatter: &UpperCaseFormatter{
				otherFormatter: &BaseFormatter{},
			},
		}

}
