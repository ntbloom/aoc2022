package days

import "os"

type Template struct {
	fd *os.File
}

func CreateTemplate(fd *os.File) *Template {
	return &Template{fd}
}

func (template *Template) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return template.solve1()
	}
	if puzzle == 2 {
		return template.solve2()
	}

	return nil
}

func (template *Template) solve1() interface{} {
	panic("implement me")
}

func (template *Template) solve2() interface{} {
	panic("implement me")
}
