package testing

import (
	"testing"
)

// Testing unit that handles comparison by itself
type ITestUnit interface {
	IsCorrect() bool
	DisplayError(t *testing.T)
}

// Test unit that holds the values for the test
type TestUnit[K any, T comparable] struct {
	Expected T
	Input    K
	Output   T
}

func (t *TestUnit[any, comparable]) IsCorrect() bool {
	return (t.Expected == t.Output)
}

func (test *TestUnit[any, comparable]) DisplayError(t *testing.T) {
	t.Errorf("Error! Input: '%v' produced '%v'. Expected '%v'\n", test.Input, test.Output, test.Expected)

}
