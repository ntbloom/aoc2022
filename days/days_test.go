package days_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ntbloom/aoc2022/parser"
	"github.com/ntbloom/aoc2022/solution"
)

func generateTest(day int, puzzle int, expected interface{}, t *testing.T) {
	filename, err := parser.GetFileName(day, parser.TestInputsDirectory)
	if err != nil {
		t.Error(err)
	}
	fd := parser.GetFileDescriptor(filename)
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Printf("Unable to close test file descriptor %s\n", fd.Name())
		}
	}(fd)

	s := solution.NewSolution(day, fd)

	actual := s.Solve(puzzle)
	if actual != expected {
		t.Errorf("wanted %d, got %d", expected, actual)
	}
}

func TestOne_Solve1(t *testing.T) {
	generateTest(1, 1, 24000, t)
}

func TestOne_Solve2(t *testing.T) {
	generateTest(1, 2, 45000, t)

}

func TestTwo_Solve1(t *testing.T) {
	generateTest(2, 1, 15, t)
}

func TestTwo_Solve2(t *testing.T) {
	generateTest(2, 2, 12, t)
}
