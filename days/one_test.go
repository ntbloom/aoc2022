package days_test

import (
	"os"
	"testing"

	"github.com/ntbloom/aoc2022/days"
	"github.com/ntbloom/aoc2022/parser"
)

func TestOne_Solve1(t *testing.T) {
	filename, err := parser.GetFileName(1, 1, parser.TestInputsDirectory)
	if err != nil {
		t.Error(err)
	}
	fd := parser.GetFileDescriptor(filename)
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)
	one := days.CreateOne(parser.GetFileDescriptor(filename))

	actual1 := one.Solve(1)
	expected1 := 24000
	if actual1 != expected1 {
		t.Errorf("wanted %d, got %d", expected1, actual1)
	}
}

func TestOne_Solve2(t *testing.T) {
	filename, err := parser.GetFileName(1, 2, parser.TestInputsDirectory)
	if err != nil {
		t.Error(err)
	}
	fd := parser.GetFileDescriptor(filename)
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)
	one := days.CreateOne(parser.GetFileDescriptor(filename))
	actual2 := one.Solve(2)
	expected2 := 45000
	if actual2 != expected2 {
		t.Errorf("wanted %d, got %d", expected2, actual2)
	}

}
