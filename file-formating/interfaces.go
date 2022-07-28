package fileformating

import "github.com/G00Z-G00Z/filename-normalizer/utils"

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
