package days

import (
	"bufio"
	"os"
)

type Six struct {
	fd *os.File
}

func CreateSix(fd *os.File) *Six {
	return &Six{fd}
}

func (six *Six) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return six.solve1()
	}
	if puzzle == 2 {
		return six.solve2()
	}

	return nil
}

func (six *Six) solve1() interface{} {
	scanner := bufio.NewScanner(six.fd)
	for scanner.Scan() {
		return FindFourInARow(scanner.Text())
	}
	panic("can't find solution")
}
func (six *Six) solve2() interface{} {
	scanner := bufio.NewScanner(six.fd)
	for scanner.Scan() {
		return FindFourteenInARow(scanner.Text())
	}
	panic("can't find solution")
}

func FindFourInARow(line string) int {
	chunk := line[:4]
	for iterator := 0; iterator < len(line)+1; iterator++ {
		group := make(map[string]bool)
		for _, char := range chunk {
			group[string(char)] = true
		}
		if len(group) == 4 {
			return iterator + 4
		}
		chunk = line[iterator+1 : iterator+5]
	}
	return -1
}
func FindFourteenInARow(line string) int {
	chunk := line[:14]
	for iterator := 0; iterator < len(line)+1; iterator++ {
		group := make(map[string]bool)
		for _, char := range chunk {
			group[string(char)] = true
		}
		if len(group) == 14 {
			return iterator + 14
		}
		chunk = line[iterator+1 : iterator+15]
	}
	return -1
}
