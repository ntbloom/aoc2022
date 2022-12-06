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
