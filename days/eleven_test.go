package days_test

import (
	"fmt"
	"testing"

	"github.com/ntbloom/aoc2022/days"
)

func TestEleven_Parse(t *testing.T) {
	fd, closed := generateTestFileDescriptor(11, 1, t)
	defer closed()
	eleven := days.CreateEleven(fd)
	fmt.Println(eleven)
}
