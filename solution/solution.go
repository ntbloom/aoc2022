package solution

import (
	"os"

	"github.com/ntbloom/aoc2022/days"
)

// Solution provides a response to the answer
type Solution interface {
	Solve(puzzle int) interface{}
}

func NewSolution(day int, fd *os.File) Solution {
	var s Solution
	switch day {
	case 1:
		s = days.CreateOne(fd)
	case 2:
		s = days.CreateTwo(fd)
	case 3:
		s = days.CreateThree(fd)
	default:
		panic("must provide valid solution day!")
	}
	return s
}
