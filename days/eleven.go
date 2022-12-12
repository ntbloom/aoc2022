package days

import (
	"os"
)

type Eleven struct {
	fd *os.File
}

func CreateEleven(fd *os.File) *Eleven {
	return &Eleven{fd}
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
	return nil
}
func (eleven *Eleven) solve2() interface{} {
	return nil
}

func (eleven *Eleven) Parse() {

}

type monkey struct {
	items       map[int]interface{}
	operation   func(new, old int) int
	criteria    func(value int) bool
	trueMonkey  *monkey
	falseMonkey *monkey
}
