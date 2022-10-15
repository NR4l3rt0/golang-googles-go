package main

import (
	"testing"
)

func TestPositiveInt(t *testing.T) {

	suite := struct {
		inputs   []int
		expected []bool
	}{
		inputs:   []int{-3, 0, 5},
		expected: []bool{false, false, true},
	}

	for i, v := range suite.inputs {
		if output := PositiveInt(v); output != suite.expected[i] {
			t.Errorf("PositiveInt(%d) - FAILED! Expected: %t, got %t", v, suite.expected[i], output)
		} else {
			t.Logf("PositiveInt(%d) - succeed, expected: %t, got %t", v, suite.expected[i], output)
		}
	}
}
