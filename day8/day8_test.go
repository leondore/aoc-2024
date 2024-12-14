package day8

import (
	"reflect"
	"testing"

	"github.com/leondore/aoc-2024/grid"
)

var test = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestNewMap(t *testing.T) {
	g := grid.Grid(test)

	want := map[rune][]grid.Coordinate{
		'0': {{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}},
		'A': {{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}},
	}
	got := NewMap(g)

	if !reflect.DeepEqual(got.Locations, want) {
		t.Errorf("NewMap() = %+v, wanted %+v", got.Locations, want)
	}
}

func TestFindAntinodes(t *testing.T) {
	g := grid.Grid(test)
	city := NewMap(g)
	list := map[grid.Coordinate]bool{}

	cases := []struct {
		coord1 grid.Coordinate
		coord2 grid.Coordinate
		want   int
	}{
		{
			grid.Coordinate{X: 8, Y: 1},
			grid.Coordinate{X: 5, Y: 2},
			4,
		},
		{
			grid.Coordinate{X: 8, Y: 1},
			grid.Coordinate{X: 4, Y: 4},
			3,
		},
		{
			grid.Coordinate{X: 8, Y: 8},
			grid.Coordinate{X: 9, Y: 9},
			12,
		},
	}

	for _, c := range cases {
		got := city.FindAntinodes(c.coord1, c.coord2, list)

		if got != c.want {
			t.Errorf("got %+v, want %+v", got, c.want)
		}
	}
}

func TestCountAntinodes(t *testing.T) {
	g := grid.Grid(test)
	city := NewMap(g)

	want := 34
	got := city.CountAntinodes()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
