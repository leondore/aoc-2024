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

	cases := []struct {
		coord1 grid.Coordinate
		coord2 grid.Coordinate
		want1  grid.Coordinate
		want2  grid.Coordinate
	}{
		{
			grid.Coordinate{X: 8, Y: 1},
			grid.Coordinate{X: 5, Y: 2},
			grid.Coordinate{X: 11, Y: 0},
			grid.Coordinate{X: 2, Y: 3},
		},
		{
			grid.Coordinate{X: 8, Y: 1},
			grid.Coordinate{X: 4, Y: 4},
			grid.Coordinate{},
			grid.Coordinate{X: 0, Y: 7},
		},
		{
			grid.Coordinate{X: 8, Y: 8},
			grid.Coordinate{X: 9, Y: 9},
			grid.Coordinate{X: 7, Y: 7},
			grid.Coordinate{X: 10, Y: 10},
		},
	}

	for _, c := range cases {
		c1, c2 := city.FindAntinodes(c.coord1, c.coord2)

		if c1 != c.want1 {
			t.Errorf("got %+v, want %+v", c1, c.want1)
		}

		if c2 != c.want2 {
			t.Errorf("got %+v, want %+v", c2, c.want2)
		}
	}
}

func TestCountAntinodes(t *testing.T) {
	g := grid.Grid(test)
	city := NewMap(g)

	want := 14
	got := city.CountAntinodes()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
