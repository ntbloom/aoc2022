package days

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Five struct {
	fd     *os.File
	stacks *Stacks
}

type Stacks struct {
	Size     int
	Elements *[]list.List
}

func NewStacks(size int, fd *os.File) *Stacks {
	scanner := bufio.NewScanner(fd)
	stacks := make([]list.List, size)
	re := regexp.MustCompile(`\[([A-Z])\[*`)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "1") {
			// goto line counter
			return nil
		}
		if strings.Contains(line, "m") {
			// goto movement parser
			return nil
		}
		matches := re.FindAllString(line, size)

		//matches := strings.Split(scanner.Text(), "  ")
		fmt.Println(matches)
	}

	return &Stacks{
		Size:     size,
		Elements: &stacks,
	}
}

func CreateFive(fd *os.File) *Five {
	return &Five{fd, nil}
}

func (five *Five) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return five.solve1()
	}
	if puzzle == 2 {
		return five.solve2()
	}

	return nil
}

func (five *Five) solve1() interface{} {
	stacks := NewStacks(3, five.fd)
	fmt.Println(stacks)
	return nil
}
func (five *Five) solve2() interface{} {
	panic("implement me")
}

func ParseCrates(line string, length int) *[]string {
	matches := make([]string, length)
	charCount := 0
	for idx, char := range line {
		if (idx % 4) == 1 {
			val := string(char)
			if val == " " {
				val = ""
			}
			matches[charCount] = val
			charCount++
		}
	}

	return &matches
}
