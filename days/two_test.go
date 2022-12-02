package days_test

import (
	"testing"

	"github.com/ntbloom/aoc2022/days"
)

type outcome struct {
	them   int
	you    int
	result int
	points int
}

func TestGetOutcomes(t *testing.T) {
	cases := []*outcome{
		// they are rock
		{
			them:   days.Rock,
			you:    days.Rock,
			result: days.RPSDraw,
			points: days.RPSRockPoints + days.RPSDrawPoints,
		},
		{
			them:   days.Rock,
			you:    days.Paper,
			result: days.RPSWin,
			points: days.RPSPaperPoints + days.RPSWinPoints,
		},
		{
			them:   days.Rock,
			you:    days.Scissors,
			result: days.RPSLose,
			points: days.RPSScissorPoints + days.RPSLosePoints,
		},

		// they are paper
		{
			them:   days.Paper,
			you:    days.Rock,
			result: days.RPSLose,
			points: days.RPSRockPoints + days.RPSLosePoints,
		},
		{
			them:   days.Paper,
			you:    days.Paper,
			result: days.RPSDraw,
			points: days.RPSPaperPoints + days.RPSDrawPoints,
		},
		{
			them:   days.Paper,
			you:    days.Scissors,
			result: days.RPSWin,
			points: days.RPSScissorPoints + days.RPSWinPoints,
		},

		// they are scissors
		{
			them:   days.Scissors,
			you:    days.Rock,
			result: days.RPSWin,
			points: days.RPSRockPoints + days.RPSWinPoints,
		},
		{
			them:   days.Scissors,
			you:    days.Paper,
			result: days.RPSLose,
			points: days.RPSPaperPoints + days.RPSLosePoints,
		},
		{
			them:   days.Scissors,
			you:    days.Scissors,
			result: days.RPSDraw,
			points: days.RPSScissorPoints + days.RPSDrawPoints,
		},
	}
	for _, v := range cases {
		result, points := days.RockPaperScissors(v.them, v.you)
		if result != v.result {
			t.Errorf("incorrect result for %d vs. %d: result=%d", v.them, v.you, result)
		}
		if points != v.points {
			t.Errorf("incorrect points for %d vs. %d: points=%d", v.them, v.you, points)
		}

	}
}

func TestDesiredOutcome(t *testing.T) {
	for _, expectedOutcome := range []int{days.RPSWin, days.RPSDraw, days.RPSLose} {
		for _, them := range []int{days.Rock, days.Paper, days.Scissors} {
			suggested := days.GetDesiredOutcome(them, expectedOutcome)
			actualResult, _ := days.RockPaperScissors(them, suggested)
			if actualResult != expectedOutcome {
				t.Errorf("Bad suggestion. actual=%d expected=%d them=%d, suggested=%d", actualResult, expectedOutcome, them, suggested)
			}

		}
	}
}
