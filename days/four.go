package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Four struct {
	fd *os.File
}

func CreateFour(fd *os.File) *Four {
	return &Four{fd}
}

func (four *Four) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return four.solve1()
	}
	if puzzle == 2 {
		return four.solve2()
	}

	return nil
}

func (four *Four) solve1() interface{} {
	fullyContained := 0
	scanner := bufio.NewScanner(four.fd)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		if len(pair) != 2 {
			panic("should only have 2")
		}

		one := getStartStop(pair[0])
		oneStart := one[0]
		oneStop := one[1]

		two := getStartStop(pair[1])
		twoStart := two[0]
		twoStop := two[1]
		if (oneStart >= twoStart && oneStop <= twoStop) || (twoStart >= oneStart && twoStop <= oneStop) {
			fullyContained++
		}
	}
	return fullyContained
}
func (four *Four) solve2() interface{} {
	scanner := bufio.NewScanner(four.fd)
	overlaps := 0
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		if len(pair) != 2 {
			panic("should only have 2")
		}

		one := getStartStop(pair[0])
		oneStart := one[0]
		oneStop := one[1]

		two := getStartStop(pair[1])
		twoStart := two[0]
		twoStop := two[1]

		if (oneStop >= twoStart && oneStop <= twoStop) || (twoStop >= oneStart && twoStop <= oneStop) {
			overlaps++
		}
	}
	return overlaps
}

func getStartStop(entry string) *[2]int {
	split := strings.Split(entry, "-")
	if len(split) != 2 {
		panic("should only have 2")
	}
	start, err := strconv.Atoi(split[0])
	if err != nil {
		panic("can't get start")
	}
	stop, err := strconv.Atoi(split[1])
	if err != nil {
		panic("can't get stop")
	}
	return &[2]int{start, stop}
}
