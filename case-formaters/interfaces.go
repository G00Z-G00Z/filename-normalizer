package caseformaters

// Function that takes a string and returns a modified string
type CaseFormatter func(s string, currentFormat CaseFormat) string

// Function that manipulates a string
type StringFormatter func(s string) string

// Function that identifies a case
type CaseIdentifier func(s string) bool

// Case format enum
type CaseFormat int

const (
	SnakeCase CaseFormat = iota
	CammelCase
	Spaces
	// Add any more cases you want to consider here
)

type caseMetaData struct {
	Case2snakeCase CaseFormatter
	Name           string
	SnakeCase2Case CaseFormatter
	IsCase         CaseIdentifier
}
