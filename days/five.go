package days

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ntbloom/aoc2022/parser"
)

const AsciiOne = 49
const AsciiNewline = 0x0A

var movementRegex *regexp.Regexp = regexp.MustCompile(`(\d+)`)

type Five struct {
	fd     *os.File
	stacks *Stacks
	size   int
}

type Stacks struct {
	Size     int
	Elements []*list.List
}

func (s *Stacks) getLastElements() *[]any {
	result := make([]any, s.Size)
	for idx, val := range s.Elements {
		last := val.Back().Value
		result[idx] = last

	}
	return &result
}

func (five *Five) newStacks(size int) *Stacks {

	scanner := bufio.NewScanner(five.fd)
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			panic(err)
		}
		newFd, err := os.Open(five.fd.Name())
		if err != nil {
			panic(err)
		}
		five.fd = newFd
	}(five.fd)

	stacks := make([]*list.List, size)
	for i := 0; i < size; i++ {
		stacks[i] = list.New()
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "1") {
			break
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
	size := getCount(fd)
	return &Five{fd, nil, size}
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
	five.stacks = five.newStacks(five.size)
	scanner := bufio.NewScanner(five.fd)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "move") {
			continue
		}
		movements := GetMovements(line)
		count := movements[0]
		popStack := movements[1]
		pushStack := movements[2]

		for i := count; i > 0; i-- {
			item := five.stacks.Elements[popStack].Back()
			five.stacks.Elements[popStack].Remove(item)
			five.stacks.Elements[pushStack].PushBack(item)
			fmt.Println(five.stacks.Elements)
		}
	}

	fmt.Println(five.stacks.getLastElements())
	return nil
}
func (five *Five) solve2() interface{} {
	panic("implement me")
}

func getCount(fd *os.File) int {
	newFd := parser.GetFileDescriptor(fd.Name())
	reader := bufio.NewReader(newFd)
	defer func() {
		err := newFd.Close()
		if err != nil {
			panic(err)
		}
	}()
	_, err := reader.ReadBytes(AsciiOne)
	if err != nil {
		panic(err)
	}
	nextLine, err := reader.ReadString(AsciiNewline)
	reader.Reset(fd)
	if err != nil {
		panic(err)
	}
	size := GetLastNumberFromString(nextLine)
	return size
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

func GetMovements(line string) []int {
	res := movementRegex.FindAllString(line, -1)
	numbers := make([]int, 3)
	if len(res) != 3 {
		log.Panicf("expected 3, got %d", len(res))
	}
	for idx, val := range res {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		// use zero-based index for the stack list
		if idx != 0 {
			num--
		}
		numbers[idx] = num
	}
	return numbers
}
