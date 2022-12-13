package days

import (
	"bufio"
	"os"
)

const Lowest uint8 = 97
const Highest uint8 = 122
const Start uint8 = 83
const End uint8 = 69

type Twelve struct {
	Elevations [][]*grid
	Start      *grid
	End        *grid
	Width      uint8
	Length     uint8
	fd         *os.File
}

func CreateTwelve(fd *os.File) *Twelve {

	twelve := Twelve{
		Start:  nil,
		Width:  0,
		Length: 0,
		fd:     fd,
	}
	twelve.parse()
	return &twelve
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
	d := bfs{
		stack:  []*grid{},
		twelve: twelve,
	}
	solution := d.find(twelve.Start)
	if !solution.equal(twelve.End) {
		panic("didn't find path!")
	}
	d.countToRoot(solution)
	return d.Count
}
func (twelve *Twelve) solve2() interface{} {
	return nil
}

func (twelve *Twelve) parse() {
	var elevations [][]*grid
	scanner := bufio.NewScanner(twelve.fd)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		twelve.Width = uint8(len(line))
		row := make([]*grid, twelve.Width)

		for idx, val := range line {
			height := uint8(val)
			node := grid{
				Height: height,
				Row:    uint8(rowCount),
				Col:    uint8(idx),
			}
			if height == Start {
				node.Height = Lowest
				twelve.Start = &node
			}
			if height == End {
				node.Height = Highest
				twelve.End = &node
			}
			row[idx] = &node
		}
		elevations = append(elevations, row)
		rowCount++
	}
	twelve.Elevations = elevations
	twelve.Length = uint8(rowCount)
}

type grid struct {
	Row       uint8
	Col       uint8
	Height    uint8
	Neighbors []*grid
	Parent    *grid
	Marked    bool
}

func (g *grid) equal(other *grid) bool {
	return g.Row == other.Row && g.Col == other.Col
}

func (twelve *Twelve) findNeighbors(g *grid) {
	var neighbors []*grid
	// check left
	if g.Col != 0 {
		left := twelve.Elevations[g.Row][g.Col-1]
		if (left.Height - g.Height) <= 1 {
			neighbors = append(neighbors, left)
		}
	}

	// check right
	if g.Col != twelve.Width-1 {
		right := twelve.Elevations[g.Row][g.Col+1]
		if (right.Height - g.Height) <= 1 {
			neighbors = append(neighbors, right)
		}
	}

	// check up
	if g.Row != 0 {
		up := twelve.Elevations[g.Row-1][g.Col]
		if (up.Height - g.Height) <= 1 {
			neighbors = append(neighbors, up)
		}
	}

	// check down
	if g.Row != twelve.Length-1 {
		down := twelve.Elevations[g.Row+1][g.Col]
		if (down.Height - g.Height) <= 1 {
			neighbors = append(neighbors, down)
		}
	}
	g.Neighbors = neighbors
}

type bfs struct {
	stack  []*grid
	twelve *Twelve
	Count  int
}

func (b *bfs) find(g *grid) *grid {
	g.Marked = true
	if g.equal(b.twelve.End) {
		return g
	}
	b.twelve.findNeighbors(g)
	for _, val := range g.Neighbors {
		if !val.Marked {
			val.Parent = g
			b.stack = append(b.stack, val)
		}
	}
	next := b.stack[0]
	b.stack = b.stack[1:]
	return b.find(next)
}

func (b *bfs) countToRoot(g *grid) {
	if g.equal(b.twelve.Start) {
		return
	}
	b.Count++
	b.countToRoot(g.Parent)
}
