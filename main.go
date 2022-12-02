package main

import (
	"fmt"
	"os"

	"github.com/ntbloom/aoc2022/days"
	"github.com/ntbloom/aoc2022/parser"
)

// Solution provides a response to the answer
type Solution interface {
	Solve(puzzle int) interface{}
}

// matchSolution runs the solution for the given day against the day and puzzle number
func matchSolution() {
	day, puzzle, fd := parser.ParseFlags()

	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)

	var solution Solution
	switch day {
	case 1:
		solution = days.CreateOne(fd)
	}

	answer := solution.Solve(puzzle)
	fmt.Printf("Solution for day %d, puzzle %d: %d", day, puzzle, answer)
}

func main() {
	matchSolution()
}
