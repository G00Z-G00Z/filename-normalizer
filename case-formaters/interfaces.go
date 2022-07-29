package caseformaters

// Function that takes a string and returns a modified string
type CaseFormatter func(string, currentFormat CaseFormat) string

// Function that takes a CaseFormatter and decorates it
type CaseFormatterDecorator func(CaseFormatter) CaseFormatter
