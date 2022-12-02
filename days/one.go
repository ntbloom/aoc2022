package days

import (
	"bufio"
	"os"
	"strconv"
)

type One struct {
	fd *os.File
}

func CreateOne(fd *os.File) *One {
	return &One{fd}
}

func (one *One) Solve(puzzle int) interface{} {
	switch puzzle {
	case 1:
		return one.solve1()
	case 2:
		return one.solve2()
	}
	return nil
}

func (one *One) solve1() interface{} {
	var max, tempSum int

	reconcile := func(temp int) {
		if temp > max {
			max = temp
		}
	}

	scanner := bufio.NewScanner(one.fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reconcile(tempSum)
			tempSum = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			tempSum += val
		}
	}
	reconcile(tempSum)
	return max
}

func (one *One) solve2() interface{} {
	var first, second, third, tempSum int

	reconcile := func(temp int) {
		if tempSum > first {
			third = second
			second = first
			first = tempSum
		} else if tempSum > second {
			third = second
			second = tempSum
		} else if tempSum > third {
			third = tempSum
		}
	}
	scanner := bufio.NewScanner(one.fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reconcile(tempSum)
			tempSum = 0

		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			tempSum += val
		}
	}
	reconcile(tempSum)
	max := first + second + third
	return max
}
