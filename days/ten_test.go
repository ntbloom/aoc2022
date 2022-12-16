package days_test

import (
	"testing"

	"github.com/ntbloom/aoc2022/days"
)

func TestPixels_GetPosition(t *testing.T) {
	pixels := days.NewPixels()
	vals := map[int][2]int{
		1:   {0, 0},
		2:   {0, 1},
		40:  {0, 39},
		41:  {1, 0},
		45:  {1, 4},
		120: {2, 39},
		160: {3, 39},
	}
	for input, rowCol := range vals {
		t.Run("row tests", func(t *testing.T) {
			actualRow, _ := pixels.GetPosition(input)
			expectedRow := rowCol[0]
			if actualRow != expectedRow {
				t.Errorf("failed to match row: wanted %d, got %d", expectedRow, actualRow)
			}
		})

		t.Run("column tests", func(t *testing.T) {
			_, actualCol := pixels.GetPosition(input)
			expectedCol := rowCol[1]
			if actualCol != expectedCol {
				t.Errorf("failed to match column: wanted %d, got %d", expectedCol, actualCol)
			}
		})
	}
}
