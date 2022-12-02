package solution_test

import (
	"strings"
	"testing"

	"github.com/ntbloom/aoc2022/parser"
	"github.com/ntbloom/aoc2022/solution"
)

// TestNewSolution makes sure we can instantiate a new solution each day
func TestNewSolution(t *testing.T) {
	days := []int{1, 2}
	for _, day := range days {
		filename, err := parser.GetFileName(day, parser.InputsDirectory)
		filename = strings.Replace(filename, "/solution", "", 1)
		if err != nil {
			panic(err)
		}
		fd := parser.GetFileDescriptor(filename)
		solution.NewSolution(day, fd)
	}
}
