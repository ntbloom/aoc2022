package days

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var itemId int = 0
var worryLevel int = 0

type Eleven struct {
	fd      *os.File
	Monkeys map[int]*monkey
}

func CreateEleven(fd *os.File) *Eleven {
	return &Eleven{fd, map[int]*monkey{}}
}

func (eleven *Eleven) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return eleven.solve1()
	}
	if puzzle == 2 {
		return eleven.solve2()
	}

	return nil
}

func (eleven *Eleven) solve1() interface{} {
	eleven.parse()
	return nil
}
func (eleven *Eleven) solve2() interface{} {
	return nil
}

func (eleven *Eleven) parse() {
	scanner := bufio.NewScanner(eleven.fd)
	var currentMonkey *monkey
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) == 1 {
			number := currentMonkey.Number
			eleven.Monkeys[number] = currentMonkey
			continue
		}
		if line[0] == "Monkey" {
			number := getNumber(line[1])
			currentMonkey = newMonkey(number)
			continue
		}
		if line[2] == "Starting" {
			for _, num := range line[4:] {
				currentMonkey.Items = append(currentMonkey.Items, newItem(getNumber(num)))
			}
			continue
		}
		if line[2] == "Operation:" {
			operator := line[6]
			element := line[7]
			currentMonkey.Operation = getOperation(operator, element)
			continue
		}
		if line[2] == "Test:" {
			currentMonkey.DivisibleBy = getNumber(line[5])
		}
		if line[5] == "true:" {
			currentMonkey.TrueMonkey = getNumber(line[9])
		}
		if line[5] == "false:" {
			currentMonkey.FalseMonkey = getNumber(line[9])
		}
	}
	number := currentMonkey.Number
	eleven.Monkeys[number] = currentMonkey
	return
	//for _, val := range line {
	//	fmt.Print(val)
	//	fmt.Print(" | ")
	//}
	//fmt.Print("\n")
}

type item struct {
	Id         int
	WorryLevel int
}

func newItem(worryLevel int) *item {
	itm := item{Id: itemId, WorryLevel: worryLevel}
	itemId++
	return &itm
}

type monkey struct {
	Number      int
	Items       []*item
	Operation   func(old, other int) int
	DivisibleBy int
	TrueMonkey  int
	FalseMonkey int
}

func newMonkey(number int) *monkey {
	return &monkey{
		Number:      number,
		Items:       []*item{},
		Operation:   nil,
		DivisibleBy: -1,
		TrueMonkey:  -1,
		FalseMonkey: -1,
	}
}

func getNumber(str string) int {
	numPattern := regexp.MustCompile(`\d+`)
	if number, err := strconv.Atoi(numPattern.FindString(str)); err != nil {
		panic(err)
	} else {
		return number
	}
}

func getOperation(operator, element string) func(int, int) int {
	if element == "old" {
		if operator == "*" {
			return func(one, _ int) int {
				return one * one
			}
		}
		if operator == "+" {
			return func(one, _ int) int {
				return one + one
			}
		}
	}
	num := getNumber(element)
	if operator == "*" {
		return func(one, _ int) int {
			return one * num
		}
	}
	if operator == "+" {
		return func(one, _ int) int {
			return one + num
		}
	}
	panic("unreachable")
}
