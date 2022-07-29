package caseformaters

// Function that takes a string and returns a modified string
type CaseFormatter func(s string, currentFormat CaseFormat) string

// Function that takes a CaseFormatter and decorates it
type CaseFormatterDecorator func(f CaseFormatter) CaseFormatter
