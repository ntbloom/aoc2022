package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ntbloom/aoc2022/elf_filesystem"
)

type Seven struct {
	fd *os.File
}

func CreateSeven(fd *os.File) *Seven {
	return &Seven{fd}
}

func (seven *Seven) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return seven.solve1()
	}
	if puzzle == 2 {
		return seven.solve2()
	}

	return nil
}

func (seven *Seven) solve1() interface{} {
	root := seven.parse()
	val := depthFirstSearch1(root, 100_000)
	return val
}
func (seven *Seven) solve2() interface{} {
	root := seven.parse()
	totalSpace := 70000000
	spaceUsed := root.TotalSize
	spaceNeeded := 30000000
	minimum := spaceNeeded - (totalSpace - spaceUsed)
	val := depthFirstSearch2(root, minimum)
	return val
}

func (seven *Seven) parse() *elf_filesystem.Dir {
	scanner := bufio.NewScanner(seven.fd)
	iterator := 0
	dir := elf_filesystem.NewDir("/", nil)
	for scanner.Scan() {
		// ignore the first line since we know it's `cd /`
		if iterator == 0 {
			iterator++
			continue
		}
		line := scanner.Text()
		lines := strings.Split(line, " ")
		switch lines[0] {
		case "$":
			if lines[1] == "cd" {
				if lines[2] == ".." {
					dir = elf_filesystem.Up(dir)
				} else {
					dir = elf_filesystem.Cd(dir, lines[2])
				}
			} else if lines[1] == "ls" {
				// we don't need to do anything with ls since the next lines will be directories
				continue
			} else {
				panic("unknown operator " + lines[1])
			}
		case "dir":
			name := lines[1]
			child := elf_filesystem.NewDir(name, dir)
			dir.Children = append(dir.Children, child)
		default:
			// it's a file
			if size, err := strconv.Atoi(lines[0]); err != nil {
				panic(err)
			} else {
				name := lines[1]
				elf_filesystem.AddSize(dir, size)
				e := elf_filesystem.NewFile(name, size)
				dir.Files = append(dir.Files, e)

			}
		}
	}
	return elf_filesystem.GetRoot(dir)
}

type depthFirstSearch struct {
	answer    int
	threshold int
	minimum   int
}

func (dfs *depthFirstSearch) calculateTotal(d *elf_filesystem.Dir) {
	d.DFSVisited = true
	for _, val := range d.Children {
		if !val.DFSVisited {
			if val.TotalSize <= dfs.threshold {
				dfs.answer += val.TotalSize
			}
		}
		dfs.calculateTotal(val)
	}
}

func (dfs *depthFirstSearch) getLeastBiggest(d *elf_filesystem.Dir) {
	d.DFSVisited = true
	for _, val := range d.Children {
		if !val.DFSVisited {
			if val.TotalSize >= dfs.minimum {
				if dfs.answer == 0 || val.TotalSize < dfs.answer {
					dfs.answer = val.TotalSize
				}
			}
		}
		dfs.getLeastBiggest(val)
	}
}

func depthFirstSearch1(root *elf_filesystem.Dir, threshold int) int {
	dfs := depthFirstSearch{answer: 0, threshold: threshold}
	dfs.calculateTotal(root)
	return dfs.answer
}

func depthFirstSearch2(root *elf_filesystem.Dir, minimum int) int {
	dfs := depthFirstSearch{answer: 0, threshold: 0, minimum: minimum}
	dfs.getLeastBiggest(root)
	return dfs.answer
}
