package days

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const AsciiOne = 49
const AsciiNewline = 0x0A

type Five struct {
	fd     *os.File
	stacks *Stacks
}

type Stacks struct {
	Size     int
	Elements []*list.List
}

func NewStacks(size int, fd *os.File) *Stacks {
	scanner := bufio.NewScanner(fd)
	reader := bufio.NewReader(fd)
	_, err := reader.ReadBytes(AsciiOne)
	if err != nil {
		panic(err)
	}
	nextLine, err := reader.ReadString(AsciiNewline)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextLine)

	stacks := make([]*list.List, size)
	for i := 0; i < size; i++ {
		stacks[i] = list.New()
	}
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
		matches := ParseCrates(line, size)
		for idx, item := range *matches {
			if item != "" {
				stacks[idx].PushBack(item)
			}
		}
	}

	return &Stacks{
		Size:     size,
		Elements: stacks,
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
	stacks := NewStacks(9, five.fd)
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

func GetLastNumberFromString(line string) int {
	re := regexp.MustCompile(`(\d+)`)
	res := re.FindAllString(line, -1)
	val, err := strconv.Atoi(res[len(res)-1])
	if err != nil {
		panic(err)
	}
	return val
}
