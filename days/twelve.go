package days

import "os"

type Twelve struct {
	fd *os.File
}

func CreateTwelve(fd *os.File) *Twelve {
	return &Twelve{fd}
}

func (twelve *Twelve) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return twelve.solve1()
	}
	if puzzle == 2 {
		return twelve.solve2()
	}

	return nil
}

func (twelve *Twelve) solve1() interface{} {
	return nil
}
func (twelve *Twelve) solve2() interface{} {
	return nil
}
