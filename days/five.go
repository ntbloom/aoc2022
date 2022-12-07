package days

import (
	"bufio"
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
	stacks map[int][]string
	size   int
}

func (five *Five) newStacks(size int) map[int][]string {

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

	stacks := make(map[int][]string, size)
	for i := 0; i < size; i++ {
		stacks[i] = []string{}
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "1") {
			break
		}
		matches := ParseCrates(line, size)
		for idx, item := range *matches {
			if item != "" {
				stacks[idx] = append(stacks[idx], item)
			}
		}
	}
	return stacks
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
		popTarget := movements[1]
		pushTarget := movements[2]

		for i := count; i > 0; i-- {
			// get the item
			item := five.stacks[popTarget][0]

			// remove it from the src
			five.stacks[movements[1]] = five.stacks[popTarget][1:]

			// add it to dest
			if len(five.stacks[pushTarget]) == 0 {
				five.stacks[pushTarget] = []string{item}
			} else {
				five.stacks[pushTarget] = append([]string{item}, five.stacks[pushTarget]...)
			}
		}
	}
	return five.getTopmost()
}

func (five *Five) getTopmost() string {
	builder := strings.Builder{}
	for i := 0; i < five.size; i++ {
		var topmost string
		target := five.stacks[i]
		switch len(target) {
		case 0:
			panic("didn't expect empty")
		case 1:
			topmost = target[0]
		default:
			topmost = target[:len(target)-1][0]
		}
		builder.WriteString(topmost)
	}
	return builder.String()
}

func (five *Five) solve2() interface{} {

	five.stacks = five.newStacks(five.size)
	scanner := bufio.NewScanner(five.fd)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "move") {
			continue
		}
		movements := GetMovements(line)
		count := movements[0]
		popTarget := movements[1]
		pushTarget := movements[2]

		// get the group of items this time
		items := five.stacks[popTarget][0:count]

		// remove it from the src
		five.stacks[popTarget] = five.stacks[popTarget][count:]

		// TODO: find out why the 2nd stack is getting written over on the first pass

		// add it to dest
		if len(five.stacks[pushTarget]) == 0 {
			five.stacks[pushTarget] = items
		} else {
			five.stacks[pushTarget] = append(items, five.stacks[pushTarget]...)
		}
	}
	return five.getTopmost()
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
