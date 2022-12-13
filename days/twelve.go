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
	Elevations *[][]*grid
	Start      *grid
	End        *grid
	Width      int
	Length     int
	fd         *os.File
}

func CreateTwelve(fd *os.File) *Twelve {

	twelve := Twelve{
		Start:  nil,
		Width:  -1,
		Length: -1,
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
	return nil
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
		twelve.Width = len(line)
		row := make([]*grid, twelve.Width)

		for idx, val := range line {
			height := uint8(val)
			node := grid{
				Height: height,
				X:      uint8(rowCount),
				Y:      uint8(idx),
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
	twelve.Elevations = &elevations
	twelve.Length = rowCount + 1
}

type grid struct {
	X         uint8
	Y         uint8
	Height    uint8
	Neighbors []*grid
	Parent    *grid
}

type bfs struct {
	root  *grid
	count int
}
