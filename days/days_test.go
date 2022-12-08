package days_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ntbloom/aoc2022/days"

	"github.com/ntbloom/aoc2022/parser"
	"github.com/ntbloom/aoc2022/solution"
)

func generateTest(day int, puzzle int, expected interface{}, t *testing.T) {
	filename, err := parser.GetFileName(day, parser.TestInputsDirectory)
	if err != nil {
		t.Error(err)
	}
	fd := parser.GetFileDescriptor(filename)
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Printf("Unable to close test file descriptor %s\n", fd.Name())
		}
	}(fd)

	s := solution.NewSolution(day, fd)

	actual := s.Solve(puzzle)
	if actual != expected {
		t.Errorf("wanted %d, got %d", expected, actual)
	}
}

func TestRegressions(t *testing.T) {
	regressions := map[int][2]interface{}{
		1: {71780, 212489},
		2: {11150, 8295},
		3: {7701, 2644},
		4: {513, 878},
		5: {"SBPQRSCDF", "RGLVRCQSB"},
		6: {1578, 2178},
		7: {1642503, 6999588},
	}
	for day, solutions := range regressions {
		for i, v := range []int{1, 2} {
			filename, err := parser.GetFileName(day, parser.InputsDirectory)
			filename = strings.Replace(filename, "days/", "", 1)
			if err != nil {
				panic(err)
			}
			fd := parser.GetFileDescriptor(filename)
			defer func(fd *os.File) {
				err = fd.Close()
				if err != nil {
					fmt.Printf("Unable to close test file descriptor %s\n", fd.Name())
				}
			}(fd)
			s := solution.NewSolution(day, fd)

			puzzle := v
			expected := solutions[i]
			actual := s.Solve(puzzle)
			if actual != expected {
				t.Errorf("day %d puzzle %d: wanted %d, got %d", day, puzzle, expected, actual)
			}
		}
	}
}

func TestOne_Solve1(t *testing.T) {
	generateTest(1, 1, 24000, t)
}

func TestOne_Solve2(t *testing.T) {
	generateTest(1, 2, 45000, t)

}

func TestTwo_Solve1(t *testing.T) {
	generateTest(2, 1, 15, t)
}

func TestTwo_Solve2(t *testing.T) {
	generateTest(2, 2, 12, t)
}

func TestThree_Solve1(t *testing.T) {
	generateTest(3, 1, 157, t)
}

func TestThree_Solve2(t *testing.T) {
	generateTest(3, 2, 70, t)
}

func TestFour_Solve1(t *testing.T) {
	generateTest(4, 1, 2, t)
}

func TestFour_Solve2(t *testing.T) {
	generateTest(4, 2, 4, t)
}

func TestFive_Solve1(t *testing.T) {
	generateTest(5, 1, "CMZ", t)
}
func TestFive_Solve2(t *testing.T) {
	generateTest(5, 2, "MCD", t)
}

func TestSix_Solve1(t *testing.T) {
	for input, expected := range map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	} {
		actual := days.FindFourInARow(input)
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	}
}
func TestSix_Solve2(t *testing.T) {
	for input, expected := range map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	} {
		actual := days.FindFourteenInARow(input)
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	}
}

func TestSeven_Solve1(t *testing.T) {
	generateTest(7, 1, 95437, t)

}

func TestSeven_Solve2(t *testing.T) {
	generateTest(7, 2, 24933642, t)
}
