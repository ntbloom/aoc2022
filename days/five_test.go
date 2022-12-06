package days_test

import (
	"reflect"
	"testing"

	"github.com/ntbloom/aoc2022/days"
)

func TestParseCrates(t *testing.T) {
	for input, expected := range map[string]*[]string{
		"[A] [B] [C]": {"A", "B", "C"},
		"    [B] [C]": {"", "B", "C"},
		"        [C]": {"", "", "C"},
		"[A]     [C]": {"A", "", "C"},
		"[A]        ": {"A", "", ""},
		"[A] [B]    ": {"A", "B", ""},
	} {
		answer := days.ParseCrates(input, 3)
		if len(*answer) != len(*expected) {
			t.Errorf("mismatched lengths: %d, %d", len(*answer), len(*expected))
		}
		if !reflect.DeepEqual(*answer, *expected) {
			t.Errorf("wanted %s, got %s", *answer, *expected)
		}
	}
}

func TestGetLastNumberFromString(t *testing.T) {

	for input, expected := range map[string]int{
		"1   2   3   4   5   6   7   8   9\n":      9,
		"1   2   3\n":                              3,
		"1   2   3   4   5   6   7   8   9   10\n": 10,
	} {
		answer := days.GetLastNumberFromString(input)
		if answer != expected {
			t.Errorf("wanted %d, got %d", expected, answer)
		}
	}
}
