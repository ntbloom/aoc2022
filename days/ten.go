package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ten struct {
	fd *os.File
}

func CreateTen(fd *os.File) *Ten {
	return &Ten{fd}
}

func (ten *Ten) Solve(puzzle int) interface{} {
	x, cycle := 1, 1
	values := map[int]int{
		20:  -1,
		60:  -1,
		100: -1,
		140: -1,
		180: -1,
		220: -1,
	}

	pixels := NewPixels()

	update := func() {
		if _, exists := values[cycle]; exists {
			values[cycle] = x
		}
	}

	scanner := bufio.NewScanner(ten.fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "noop" {
			update()
			pixels.Draw(cycle, x)
			cycle++
			continue
		} else {
			if num, err := strconv.Atoi(strings.Split(line, " ")[1]); err != nil {
				panic(err)
			} else {
				update()
				pixels.Draw(cycle, x)
				cycle++

				update()
				pixels.Draw(cycle, x)
				x += num
				cycle++
				continue
			}
		}
	}
	var signalStrength int
	for key, value := range values {
		signalStrength += key * value
	}

	pixels.Print()
	return signalStrength
}

type Pixels struct {
	grid [6][40]string
}

func NewPixels() *Pixels {
	grid := [6][40]string{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			grid[i][j] = "?"
		}
	}
	return &Pixels{grid: grid}
}

func (p *Pixels) GetPosition(cycle int) (row, col int) {
	row = (cycle - 1) / 40
	col = (cycle - 1) % 40
	return
}

func (p *Pixels) Draw(cycle, xPos int) {
	row, col := p.GetPosition(cycle)
	var pixel string
	if xPos == col || xPos == col-1 || xPos == col+1 {
		pixel = "#"
	} else {
		pixel = "."
	}
	if cycle <= 240 {
		p.grid[row][col] = pixel

	}
}

func (p *Pixels) Print() {
	fmt.Println()
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Printf("%s", p.grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}
