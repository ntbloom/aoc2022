package days

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

const Lowest int = 97
const Highest int = 122
const Start int = 83
const End int = 69

type Twelve struct {
	Elevations [][]*grid
	Start      *grid
	End        *grid
	Width      int
	Length     int
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
	a := newAstar(twelve)
	defer func() {
		if r := recover(); r != nil {
			twelve.print()
		} else {
			twelve.print()
		}
	}()
	solution := a.find(twelve.Start)
	if !solution.equal(twelve.End) {
		panic("didn't find path!")
	}
	a.countToRoot(solution)
	return a.Count
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
		twelve.Width = int(len(line))
		row := make([]*grid, twelve.Width)

		for idx, val := range line {
			height := int(val)
			node := grid{
				Letter: val,
				Height: height,
				Row:    int(rowCount),
				Col:    int(idx),
			}
			if height == Start {
				node.Height = Lowest
				node.Start = true
				twelve.Start = &node
			}
			if height == End {
				node.Height = Highest
				node.Finish = true
				twelve.End = &node
			}
			row[idx] = &node
		}
		elevations = append(elevations, row)
		rowCount++
	}
	twelve.Elevations = elevations
	twelve.Length = int(rowCount)
}

func (twelve *Twelve) print() {
	for idx, line := range twelve.Elevations {
		fmt.Printf("%d: ", idx)
		if idx < 10 {
			fmt.Printf(" ")
		}
		for _, char := range line {
			if char.Start {
				fmt.Printf("S")
				continue
			}
			if char.Finish {
				fmt.Printf("E")
				continue
			}
			if char.OnThePath {
				fmt.Printf("\u001b[31m%s\033[0m", string(char.Letter))
				continue
			}
			//if char.Marked {
			//	fmt.Printf("+")
			//	continue
			//}
			fmt.Printf("%s", string(char.Letter))

		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

type grid struct {
	Letter    rune
	Row       int
	Col       int
	Height    int
	Neighbors []*grid
	Parent    *grid
	Marked    bool
	Priority  int
	Index     int
	OnThePath bool
	Start     bool
	Finish    bool
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

type priorityQueue struct {
	queue []*grid
}

func (pq *priorityQueue) Len() int {
	return len(pq.queue)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return pq.queue[i].Priority < pq.queue[j].Priority
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].Index = i
	pq.queue[j].Index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(pq.queue)
	item := x.(*grid)
	item.Index = n
	pq.queue = append(pq.queue, item)
}

func (pq *priorityQueue) Pop() any {
	old := pq.queue
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	pq.queue = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(grid *grid, priority int) {
	grid.Priority = priority
	heap.Fix(pq, grid.Index)
}

type aStar struct {
	pq     *priorityQueue
	twelve *Twelve
	Count  int
}

func newAstar(twelve *Twelve) *aStar {
	queue := make([]*grid, 0)
	pq := &priorityQueue{queue: queue}
	pq.Push(twelve.Start)
	return &aStar{
		pq:     pq,
		twelve: twelve,
		Count:  0,
	}
}

func (twelve *Twelve) heuristic(g *grid) int {
	rowDist := math.Abs(float64(g.Row) - float64(twelve.End.Row))
	colDist := math.Abs(float64(g.Col) - float64(twelve.End.Col))
	return int(rowDist + colDist)
}

func (a *aStar) find(g *grid) *grid {
	g.Marked = true
	if g.equal(a.twelve.End) {
		return g
	}
	a.twelve.findNeighbors(g)
	for _, val := range g.Neighbors {
		heuristic := a.twelve.heuristic(val)
		if !val.Marked { //|| a.twelve.heuristic(val) < a.twelve.heuristic(g) {
			val.Parent = g
			a.pq.Push(val)
			a.pq.update(val, heuristic)
		}
	}
	next := a.pq.Pop().(*grid)
	return a.find(next)
}

func (a *aStar) countToRoot(g *grid) {
	if !g.Start || !g.Finish {
		g.OnThePath = true
	}
	if g.equal(a.twelve.Start) {
		return
	}
	a.Count++
	a.countToRoot(g.Parent)
}
