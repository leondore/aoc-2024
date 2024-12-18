package day10

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/leondore/aoc-2024/grid"
	"github.com/leondore/aoc-2024/utils"
)

func TestGetCardinals(t *testing.T) {
	cases := []struct {
		head grid.Coordinate
		want [4]grid.Coordinate
	}{
		{
			grid.Coordinate{X: 2, Y: 0},
			[4]grid.Coordinate{{X: 2, Y: -1}, {X: 3, Y: 0}, {X: 2, Y: 1}, {X: 1, Y: 0}},
		},
		{
			grid.Coordinate{X: 4, Y: 2},
			[4]grid.Coordinate{{X: 4, Y: 1}, {X: 5, Y: 2}, {X: 4, Y: 3}, {X: 3, Y: 2}},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("[%d,%d]", c.head.X, c.head.Y), func(t *testing.T) {
			got := getCardinals(c.head)

			if got != c.want {
				t.Errorf("expected %v to produce %v, but got %v", c.head, c.want, got)
			}
		})
	}
}

func TestCountTrails(t *testing.T) {
	cases := []struct {
		head grid.Coordinate
		want int32
	}{
		{grid.Coordinate{X: 2, Y: 0}, 20},
		{grid.Coordinate{X: 4, Y: 0}, 24},
		{grid.Coordinate{X: 6, Y: 6}, 8},
	}

	test, _ := utils.ProcessInput("./test.txt")
	g := grid.Grid(test)

	for _, c := range cases {
		t.Run(fmt.Sprintf("[%d,%d]", c.head.X, c.head.Y), func(t *testing.T) {
			var score atomic.Int32
			countTrails(g, c.head, &score)

			if score.Load() != c.want {
				t.Errorf("expected %v to produce %d trails, but got %d", c.head, c.want, score.Load())
			}
		})
	}
}

func TestDay10(t *testing.T) {
	got, _ := Day10("./test.txt")
	want := 81

	if got != want {
		t.Errorf("expected score of %d, but got %d", want, got)
	}
}
