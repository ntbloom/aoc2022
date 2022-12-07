package main

import (
	"fmt"
	"os"

	"github.com/ntbloom/aoc2022/solution"

	"github.com/ntbloom/aoc2022/parser"
)

func main() {
	day, puzzle, fd := parser.ParseFlags()
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Printf("Unable to close file descriptor %s\n", fd.Name())
		}
	}(fd)

	solver := solution.NewSolution(day, fd)
	answer := solver.Solve(puzzle)
	fmt.Printf("Solution for day %d, puzzle %d: %d\n", day, puzzle, answer)

}
