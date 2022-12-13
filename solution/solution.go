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
	case 4:
		s = days.CreateFour(fd)
	case 5:
		s = days.CreateFive(fd)
	case 6:
		s = days.CreateSix(fd)
	case 7:
		s = days.CreateSeven(fd)
	case 8:
		s = days.CreateEight(fd)
	case 9:
		s = days.CreateNine(fd)
	case 12:
		s = days.CreateTwelve(fd)
	default:
		panic("must provide valid solution day!")
	}
	return s
}
