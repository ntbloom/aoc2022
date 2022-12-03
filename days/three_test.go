package days_test

import (
	"sort"
	"testing"

	"github.com/ntbloom/aoc2022/days"
)

func TestSplitItem(t *testing.T) {
	for input, result := range map[string][]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp":         {"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": {"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		"PmmdzqPrVvPwwTWBwg":               {"PmmdzqPrV", "vPwwTWBwg"},
	} {
		expFirst, expSecond := result[0], result[1]
		actualFirst, actualSecond := days.SplitItem(input)
		if actualFirst != expFirst || actualSecond != expSecond {
			t.Errorf("Expected (%s, %s), got (%s, %s)", expFirst, expSecond, actualFirst, actualSecond)
		}
	}
}

func TestFindCommonsInTwo(t *testing.T) {
	for _, vals := range map[int][][]string{
		1: {{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, {"p"}},
		2: {{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, {"L"}},
		3: {{"PmmdzqPrV", "vPwwTWBwg"}, {"P"}},
		4: {{"PmmdzqPrV", "vPmwTWBwg"}, {"P", "m"}},
		5: {{"PqmmdzqPrV", "qvPmwTWBwg"}, {"P", "m", "q"}},
	} {
		var first, second string

		first = vals[0][0]
		second = vals[0][1]
		expected := vals[1]
		actual := *days.FindCommonsInTwo(first, second)
		sort.Strings(actual)
		sort.Strings(expected)
		if len(actual) != len(expected) {
			t.Errorf("wrong number")
		}
		for i, v := range actual {
			if v != expected[i] {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		}

	}

}

func TestFindCommonsInThree(t *testing.T) {
	for _, vals := range map[int][]string{
		1: {"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "r"},
		2: {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw", "Z"},
	} {
		var first, second string

		first = vals[0]
		second = vals[1]
		third := vals[2]
		expected := vals[3]
		actual := *days.FindCommonsInThree(first, second, third)
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	}
}
