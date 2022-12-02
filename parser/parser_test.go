package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ntbloom/aoc2022/parser"
)

func TestGetFileName(t *testing.T) {
	day := 24
	puzzle := 3
	actual, err := parser.GetFileName(day, puzzle, parser.TestInputsDirectory)
	if err != nil {
		t.Error(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("%s/test_inputs/%d-%d", dir, day, puzzle)
	if expected != actual {
		t.Errorf("wanted %s, got %s", expected, actual)
	}

}
