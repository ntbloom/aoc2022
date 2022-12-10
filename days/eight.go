package days

import (
	"bufio"
	"os"
	"strconv"

	"github.com/ntbloom/aoc2022/parser"
)

type Eight struct {
	fd     *os.File
	forest *[][]*TreeNode
	width  int
	length int
}

func CreateEight(fd *os.File) *Eight {
	length, width := getCounts(fd)
	trees := makeTree(length, width)
	scanner := bufio.NewScanner(fd)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for idx, val := range line {
			if height, err := strconv.Atoi(string(val)); err != nil {
				panic(err)
			} else {
				node := TreeNode{
					Height: height,
					Row:    i,
					Column: idx,
				}
				(*trees)[i][idx] = &node
			}
		}
		i++
	}
	return &Eight{fd, trees, width, length}
}

func (eight *Eight) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return eight.solve1()
	}
	if puzzle == 2 {
		return eight.solve2()
	}

	return nil
}

func (eight *Eight) solve1() interface{} {
	return eight.findVisibles()
}
func (eight *Eight) solve2() interface{} {
	maxScore := 0
	for _, trees := range *eight.forest {
		for _, tree := range trees {
			eight.checkEast(tree)
			eight.checkWest(tree)
			eight.checkNorth(tree)
			eight.checkSouth(tree)
			score := tree.EastDistance * tree.WestDistance * tree.NorthDistance * tree.SouthDistance
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func makeTree(length, width int) *[][]*TreeNode {
	rows := make([][]*TreeNode, length)
	for i := 0; i < length; i++ {
		rows[i] = make([]*TreeNode, width)
	}
	return &rows
}
func getCounts(fd *os.File) (length int, width int) {
	newFd := parser.GetFileDescriptor(fd.Name())
	defer func(newFd *os.File) {
		err := newFd.Close()
		if err != nil {
			panic(err)
		}
	}(newFd)

	firstScanner := bufio.NewScanner(newFd)
	i := 0
	for firstScanner.Scan() {
		if i == 0 {
			width = len(firstScanner.Text())
		}
		i++
	}
	length = i
	return
}

type TreeNode struct {
	Height        int
	Row           int
	Column        int
	NorthDistance int
	SouthDistance int
	WestDistance  int
	EastDistance  int
}

func (eight *Eight) findVisibles() int {
	visibles := 0
	for _, trees := range *eight.forest {
		for _, tree := range trees {
			if eight.checkEast(tree) || eight.checkWest(tree) || eight.checkNorth(tree) || eight.checkSouth(tree) {
				visibles++
			}
		}
	}
	return visibles
}

func (eight *Eight) checkNorth(t *TreeNode) (visible bool) {
	if t.Row == 0 {
		visible = true
		t.NorthDistance = 0
	} else {
		distance := 0
		visible = true
		for i := t.Row; i > 0; i-- {
			northernHeight := (*eight.forest)[i-1][t.Column].Height
			distance++
			if northernHeight >= t.Height {
				visible = false
				break
			}
		}
		t.NorthDistance = distance
	}
	return
}

func (eight *Eight) checkSouth(t *TreeNode) (visible bool) {
	if t.Row == eight.length-1 {
		visible = true
		t.SouthDistance = 0
	} else {
		distance := 0
		visible = true
		for i := t.Row; i < eight.length-1; i++ {
			distance++
			southernHeight := (*eight.forest)[i+1][t.Column].Height
			if southernHeight >= t.Height {
				visible = false
				break
			}
		}
		t.SouthDistance = distance
	}
	return
}

func (eight *Eight) checkWest(t *TreeNode) (visible bool) {
	if t.Column == 0 {
		visible = true
		t.WestDistance = 0
	} else {
		distance := 0
		visible = true
		for i := t.Column; i > 0; i-- {
			distance++
			westernHeight := (*eight.forest)[t.Row][i-1].Height
			if westernHeight >= t.Height {
				visible = false
				break
			}
		}
		t.WestDistance = distance
	}
	return
}

func (eight *Eight) checkEast(t *TreeNode) (visible bool) {
	if t.Column == eight.width-1 {
		visible = true
		t.EastDistance = 0
	} else {
		distance := 0
		visible = true
		for i := t.Column; i < eight.width-1; i++ {
			distance++
			easternHeight := (*eight.forest)[t.Row][i+1].Height
			if easternHeight >= t.Height {
				visible = false
				break
			}
		}
		t.EastDistance = distance
	}
	return
}
