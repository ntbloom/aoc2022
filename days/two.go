package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Two struct {
	fd *os.File
}

const (
	Rock     = 0
	Paper    = 1
	Scissors = 2

	RPSRockPoints    = 1
	RPSPaperPoints   = 2
	RPSScissorPoints = 3
	RPSWinPoints     = 6
	RPSDrawPoints    = 3
	RPSLosePoints    = 0

	RPSWin  = 0
	RPSLose = 1
	RPSDraw = 2
)

func CreateTwo(fd *os.File) *Two {
	return &Two{fd}
}

func (two *Two) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return two.solve1()
	}
	if puzzle == 2 {
		return two.solve2()
	}

	return nil
}

func (two *Two) solve1() int {
	firstMap := map[string]int{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	secondMap := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	var totalScore int
	scanner := bufio.NewScanner(two.fd)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")
		them := firstMap[res[0]]
		you := secondMap[res[1]]
		_, points := RockPaperScissors(them, you)
		totalScore += points
	}
	return totalScore
}

func (two *Two) solve2() int {

	firstMap := map[string]int{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	secondMap := map[string]int{
		"X": RPSLose,
		"Y": RPSDraw,
		"Z": RPSWin,
	}

	var totalScore int
	scanner := bufio.NewScanner(two.fd)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " ")
		them := firstMap[res[0]]
		you := GetDesiredOutcome(them, secondMap[res[1]])
		_, points := RockPaperScissors(them, you)
		totalScore += points
	}
	return totalScore
}

func RockPaperScissors(them, you int) (result int, points int) {
	msg := fmt.Sprintf("invalid outcome: them=%d, you=%d", them, you)
	switch them {
	case Rock:
		switch you {
		case Rock:
			result = RPSDraw
		case Paper:
			result = RPSWin
		case Scissors:
			result = RPSLose
		default:
			panic(msg)
		}
	case Paper:
		switch you {
		case Rock:
			result = RPSLose
		case Paper:
			result = RPSDraw
		case Scissors:
			result = RPSWin
		default:
			panic(msg)
		}

	case Scissors:
		switch you {
		case Rock:
			result = RPSWin
		case Paper:
			result = RPSLose
		case Scissors:
			result = RPSDraw
		default:
			panic(msg)
		}
	}
	switch you {
	case Rock:
		points = RPSRockPoints
	case Paper:
		points = RPSPaperPoints
	case Scissors:
		points = RPSScissorPoints
	}
	switch result {
	case RPSWin:
		points += RPSWinPoints
	case RPSLose:
		points += RPSLosePoints
	case RPSDraw:
		points += RPSDrawPoints
	}
	return
}

func GetDesiredOutcome(them, expectedResult int) int {
	switch expectedResult {
	case RPSWin:
		switch them {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		default:
			panic("unreachable!")
		}
	case RPSDraw:
		return them
	case RPSLose:
		switch them {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		default:
			panic("unreachable!")
		}
	}
	panic("unreachable")
}
