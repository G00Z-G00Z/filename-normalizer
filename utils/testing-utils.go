package utils

import (
	"testing"
)

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

func (test *TestUnit[any, comparable]) SetOutput(out comparable) {
	test.Output = out

}
