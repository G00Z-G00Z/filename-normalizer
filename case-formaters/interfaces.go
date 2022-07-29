package caseformaters

// Function that takes a string and returns a modified string
type CaseFormatter func(s string, currentFormat CaseFormat) string

// Function that manipulates a string
type StringFormatter func(s string) string

// Function that identifies a case
type CaseIdentifier func(s string) bool

// Case format enum
type CaseFormat int

// You can add here any other case you want to use
const (
	SnakeCase CaseFormat = iota
	CammelCase
	SpacesCase
	Unkwon
	// Add any more cases you want to consider here
)

// Metadata for adding more cases
type caseMetaData struct {
	Case2snakeCase CaseFormatter
	Name           string
	SnakeCase2Case StringFormatter
	IsCase         CaseIdentifier
}

// Map with all the metadata to handle caseFormatters
type caseFormatterMetaDataMap map[CaseFormat]caseMetaData
