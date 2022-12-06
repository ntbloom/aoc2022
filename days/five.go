package days

import "os"

type Five struct {
	fd *os.File
}

func CreateFive(fd *os.File) *Five {
	return &Five{fd}
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
	panic("implement me")
}
func (five *Five) solve2() interface{} {
	panic("implement me")
}
