package days

import (
	"bufio"
	"fmt"
	"math/big"
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
	eleven := Eleven{fd, []*monkey{}}
	eleven.parse()
	return &eleven
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
	rounds := 20
	for i := 0; i < rounds; i++ {
		for _, m := range eleven.Monkeys {
			for _, itm := range m.Items {
				m.inspect(&itm, 3, eleven)
			}
			m.Items = []item{}
		}

	}
	return eleven.calculateMonkeyBusiness()
}
func (eleven *Eleven) solve2() interface{} {
	rounds := 10000
	for i := 0; i < rounds; i++ {
		for _, m := range eleven.Monkeys {
			for _, itm := range m.Items {
				m.inspect(&itm, 1, eleven)
			}
			m.Items = []item{}
		}

	}
	return eleven.calculateMonkeyBusiness()
}

func (eleven *Eleven) printMonkeys() {
	for idx, mky := range eleven.Monkeys {
		fmt.Printf("monkey%d ", idx)
		for _, val := range mky.Items {
			fmt.Printf("%d ", val.WorryLevel)
		}
		fmt.Println()
	}
}

func (eleven *Eleven) calculateMonkeyBusiness() int {
	var inspectionCounts []int
	for _, m := range eleven.Monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCount)
	}
	sort.Ints(inspectionCounts)
	max, penultimate := inspectionCounts[len(inspectionCounts)-1], inspectionCounts[(len(inspectionCounts)-2)]
	monkeyBusiness := max * penultimate
	return monkeyBusiness
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
				currentMonkey.Items = append(currentMonkey.Items, *newItem(big.NewInt(int64(getNumber(num)))))
			}
			continue
		}
		if line[2] == "Operation:" {
			currentMonkey.Operation = line[6]
			currentMonkey.OperationElement = line[7]
			continue
		}
		if line[2] == "Test:" {
			currentMonkey.DivisibleBy = big.NewInt(int64(getNumber(line[5])))
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
	WorryLevel *big.Int
}

func newItem(worryLevel *big.Int) *item {
	itm := item{Id: itemId, WorryLevel: worryLevel}
	itemId++
	return &itm
}

type monkey struct {
	Number           int
	Items            []item
	Operation        string
	OperationElement string
	DivisibleBy      *big.Int
	TrueMonkey       int
	FalseMonkey      int
	InspectionCount  int
}

func newMonkey(number int) *monkey {
	return &monkey{
		Number:           number,
		Items:            []item{},
		Operation:        "",
		OperationElement: "",
		DivisibleBy:      big.NewInt(-1),
		TrueMonkey:       -1,
		FalseMonkey:      -1,
		InspectionCount:  0,
	}
}

func (m *monkey) inspect(i *item, divisor int, eleven *Eleven) {
	m.InspectionCount++
	i.WorryLevel = m.operate(i.WorryLevel)
	i.WorryLevel.Div(i.WorryLevel, big.NewInt(int64(divisor)))

	res := big.NewInt(0)
	res.Div(i.WorryLevel, m.DivisibleBy)
	pass := res.Cmp(big.NewInt(0)) == 0

	if pass {
		eleven.Monkeys[m.TrueMonkey].Items = append(eleven.Monkeys[m.TrueMonkey].Items, *i)
	} else {
		eleven.Monkeys[m.FalseMonkey].Items = append(eleven.Monkeys[m.FalseMonkey].Items, *i)
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

func (m *monkey) operate(worry *big.Int) *big.Int {
	if m.OperationElement == "old" {
		if m.Operation == "*" {
			return worry.Mul(worry, worry)
		}
		if m.Operation == "+" {
			return worry.Add(worry, worry)
		}
	}

	num := big.NewInt(int64(getNumber(m.OperationElement)))
	if m.Operation == "*" {
		return worry.Mul(worry, num)
	}
	if m.Operation == "+" {
		return worry.Add(worry, num)
	}
	panic("unreachable")
}
