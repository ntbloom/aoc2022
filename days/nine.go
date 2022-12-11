package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Nine struct {
	fd *os.File
}

func CreateNine(fd *os.File) *Nine {
	return &Nine{fd}
}

func (nine *Nine) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return nine.solve1()
	}
	if puzzle == 2 {
		return nine.solve2()
	}

	return nil
}

func (nine *Nine) solve1() interface{} {
	head := newRopeHead()
	tail := newRopeTail(head.Position)

	scanner := bufio.NewScanner(nine.fd)
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")
		letter := splits[0]
		count, err := strconv.Atoi(splits[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			switch letter {
			case "L":
				head.Position.Left()
			case "R":
				head.Position.Right()
			case "U":
				head.Position.Up()
			case "D":
				head.Position.Down()
			}
			tail.follow()
		}
	}
	return len(tail.Visited)

}
func (nine *Nine) solve2() interface{} {
	scanner := bufio.NewScanner(nine.fd)
	head := newRopeHead()
	tails := [9]*ropeTail{}
	currentTail := newRopeTail(head.Position)
	tails[0] = currentTail
	for i := 1; i < 9; i++ {
		t := newRopeTail(currentTail.Position)
		tails[i] = t
		currentTail = t
	}

	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")
		letter := splits[0]
		count, err := strconv.Atoi(splits[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			switch letter {
			case "L":
				head.Position.Left()
			case "R":
				head.Position.Right()
			case "U":
				head.Position.Up()
			case "D":
				head.Position.Down()
			}
			for k := 0; k < len(tails); k++ {
				tails[k].follow()
			}
		}
	}
	return len(tails[len(tails)-1].Visited)
}

/* head and tail */

type ropeHead struct {
	Position *position
}

func newRopeHead() *ropeHead {
	return &ropeHead{
		Position: &position{0, 0},
	}
}

type ropeTail struct {
	Head     *position
	Position *position
	Visited  map[string]bool
}

func newRopeTail(head *position) *ropeTail {
	return &ropeTail{
		Head:     head,
		Position: &position{0, 0},
		Visited:  map[string]bool{head.toString(): true},
	}
}

func (rt *ropeTail) distanceToHead() (int, int) {
	diffX := rt.Head.X - rt.Position.X
	diffY := rt.Head.Y - rt.Position.Y
	return diffX, diffY
}

func (rt *ropeTail) nextToHead() bool {
	diffX, diffY := rt.distanceToHead()
	if (diffX == -1 || diffX == 0 || diffX == 1) && (diffY == -1 || diffY == 0 || diffY == 1) {
		return true
	}
	return false
}

func (rt *ropeTail) follow() {
	if rt.nextToHead() {
		return
	}

	diffX, diffY := rt.distanceToHead()
	switch diffX {
	case -2:
		switch diffY {
		case 2:
			rt.Position.Left()
			rt.Position.Up()
		case 1:
			rt.Position.Left()
			rt.Position.Up()
		case 0:
			rt.Position.Left()
		case -1:
			rt.Position.Left()
			rt.Position.Down()
		case -2:
			rt.Position.Left()
			rt.Position.Down()
		default:
			panic("unreachable")
		}
	case -1:
		switch diffY {
		case 2:
			rt.Position.Left()
			rt.Position.Up()
		case -2:
			rt.Position.Left()
			rt.Position.Down()
		default:
			panic("unreachable")
		}
	case 0:
		switch diffY {
		case 2:
			rt.Position.Up()
		case -2:
			rt.Position.Down()
		default:
			panic("unreachable")
		}
	case 1:
		switch diffY {
		case 2:
			rt.Position.Right()
			rt.Position.Up()
		case -2:
			rt.Position.Right()
			rt.Position.Down()
		default:
			panic("unreachable")
		}
	case 2:
		switch diffY {
		case 2:
			rt.Position.Right()
			rt.Position.Up()
		case 1:
			rt.Position.Right()
			rt.Position.Up()
		case 0:
			rt.Position.Right()
		case -1:
			rt.Position.Right()
			rt.Position.Down()
		case -2:
			rt.Position.Right()
			rt.Position.Down()
		default:
			panic("unreachable")
		}
	}
	// mark it as visited
	rt.Visited[rt.Position.toString()] = true

}

type position struct {
	X int
	Y int
}

func (p *position) toString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p *position) equal(other *position) bool {
	res := p.X == other.X && p.Y == other.Y
	return res
}
func (p *position) Left() {
	p.X--
}

func (p *position) Right() {
	p.X++
}

func (p *position) Up() {
	p.Y++
}

func (p *position) Down() {
	p.Y--
}
