package days

import (
	"bufio"
	"fmt"
	"os"
)

type Three struct {
	fd *os.File
}

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func CreateThree(fd *os.File) *Three {
	return &Three{fd}
}

func (three *Three) Solve(puzzle int) interface{} {
	if puzzle == 1 {
		return three.solve1()
	}
	if puzzle == 2 {
		return three.solve2()
	}

	return nil
}

func (three *Three) solve1() interface{} {
	scanner := bufio.NewScanner(three.fd)
	priority := 0
	for scanner.Scan() {
		for _, item := range *FindCommonsInTwo(SplitItem(scanner.Text())) {
			priority += priorities[item]
		}
	}
	return priority
}

func (three *Three) solve2() interface{} {
	scanner := bufio.NewScanner(three.fd)
	buffer := make([]string, 3)
	counter := 0
	priority := 0
	groups := 0
	var last, penultimate, penpenultimate string
	for scanner.Scan() {
		buffer[counter] = scanner.Text()

		if counter == 2 {
			common := *FindCommonsInThree(buffer[0], buffer[1], buffer[2])
			addition := priorities[common]
			if addition < 1 || addition > 52 {
				panic("invalid value")
			}
			if common == "" {
				panic("can't find common element")
			}
			priority += addition
			penpenultimate = buffer[0]
			penultimate = buffer[1]
			last = buffer[2]
			buffer[0] = ""
			buffer[1] = ""
			buffer[2] = ""
			counter = 0
			groups += 1
			continue
		}
		counter++
	}
	fmt.Println(last, penultimate, penpenultimate, groups)
	if counter != 0 {
		panic("miscounted!")
	}
	return priority

}

func SplitItem(item string) (string, string) {
	length := len(item) / 2
	first := item[0:length]
	second := item[length:]
	return first, second
}

func FindCommonsInTwo(str1, str2 string) *[]string {
	maps := make(map[string]bool, len(str1))
	common := 0
	exists := func(val string, arr map[string]bool) bool {
		if _, in := arr[val]; in {
			return true
		} else {
			return false
		}
	}
	for _, v := range str1 {
		maps[string(v)] = false
	}
	for _, v := range str2 {
		asStr := string(v)
		if exists(asStr, maps) && maps[asStr] != true {
			maps[asStr] = true
			common += 1
		}
	}
	result := make([]string, common)
	i := 0
	for key, val := range maps {
		if val == true {
			result[i] = key
			i++
		}

	}
	return &result
}

func FindCommonsInThree(str1, str2, str3 string) *string {
	if str1 == "" || str2 == "" || str3 == "" {
		panic("received empty string")
	}
	maps := make(map[string]bool, len(str1))
	exists := func(val string, arr map[string]bool) bool {
		if _, in := arr[val]; in {
			return true
		} else {
			return false
		}
	}
	for _, v := range str1 {
		maps[string(v)] = false
	}
	for _, v := range str2 {
		asStr := string(v)
		if exists(asStr, maps) {
			maps[asStr] = true
		}
	}

	for _, v := range str3 {
		asStr := string(v)
		if exists(asStr, maps) && maps[asStr] == true {
			common := &asStr
			if len(*common) != 1 {
				panic("should only be a single char")
			}
			return common
		}
	}
	panic("should be unreachable!")
}
