package day4

import (
	"testing"
)

func TestDay4(t *testing.T) {
	testPath := "./test.txt"

	got, err := Day4(testPath)
	want := 18

	if err != nil {
		t.Errorf("Day4() failed with error %v", err)
	}

	if got != want {
		t.Errorf("Day4() = %v; want %v", got, want)
	}
}

func TestCheckPossiblePath(t *testing.T) {
	test := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	grid := Grid(test)

	cases := []struct {
		coord    Coordinate
		dir      Coordinate
		expected bool
		name     string
	}{
		{Coordinate{0, 0}, Coordinate{0, -1}, false, "[0, 0] -> North"},
		{Coordinate{6, 4}, Coordinate{0, -1}, true, "[6, 4] -> North"},
		{Coordinate{6, 4}, Coordinate{-1, 0}, true, "[6, 4] -> West"},
		{Coordinate{5, 9}, Coordinate{0, 1}, false, "[5, 9] -> South"},
		{Coordinate{5, 9}, Coordinate{1, -1}, true, "[5, 9] -> NorthEast"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := checkPossiblePath(&grid, c.coord, c.dir, 0)

			if got != c.expected {
				t.Errorf("checkPossiblePath() = %v; want true", got)
			}
		})
	}
}
