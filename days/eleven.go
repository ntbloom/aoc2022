package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var itemId int = 0

type Eleven struct {
	fd      *os.File
	Monkeys []*monkey
}

func CreateEleven(fd *os.File) *Eleven {
	return &Eleven{fd, []*monkey{}}
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
	rounds := 20
	for i := 0; i < rounds; i++ {
		for _, m := range eleven.Monkeys {
			for _, itm := range m.Items {
				m.inspect(itm)
			}
		}
	}
	var inspectionCounts []int
	for _, m := range eleven.Monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCount)
	}
	sort.Ints(inspectionCounts)
	max, penultimate := inspectionCounts[len(inspectionCounts)-1], inspectionCounts[(len(inspectionCounts)-2)]
	monkeyBusiness := max * penultimate
	return monkeyBusiness
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
			eleven.Monkeys = append(eleven.Monkeys, currentMonkey)
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
	eleven.Monkeys = append(eleven.Monkeys, currentMonkey)
	return
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
	Number          int
	Items           []*item
	Operation       func(old, other int) int
	DivisibleBy     int
	TrueMonkey      int
	FalseMonkey     int
	InspectionCount int
}

func newMonkey(number int) *monkey {
	return &monkey{
		Number:          number,
		Items:           []*item{},
		Operation:       nil,
		DivisibleBy:     -1,
		TrueMonkey:      -1,
		FalseMonkey:     -1,
		InspectionCount: 0,
	}
}

func (m *monkey) inspect(i *item) {
	m.InspectionCount++
	fmt.Println(i)
	//panic("implement me")
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
