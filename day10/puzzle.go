package day10

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/leondore/aoc-2024/grid"
	"github.com/leondore/aoc-2024/utils"
)

const (
	trailHead = byte('0')
	end       = 9
)

var (
	N = grid.Coordinate{X: 0, Y: -1}
	E = grid.Coordinate{X: 1, Y: 0}
	S = grid.Coordinate{X: 0, Y: 1}
	W = grid.Coordinate{X: -1, Y: 0}
)

func Day10(path string) (int, error) {
	input, err := utils.ProcessInput(path)
	if err != nil {
		return 0, fmt.Errorf("could not process file: %w", err)
	}
	g := grid.Grid(input)

	var score atomic.Int32
	var wg sync.WaitGroup

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			coord := grid.Coordinate{X: x, Y: y}
			height := g.Get(coord)

			if height == trailHead {
				wg.Add(1)
				go func() {
					countTrails(g, coord, &score)
					wg.Done()
				}()
			}
		}
	}

	wg.Wait()

	return int(score.Load()), nil
}

func getCardinals(c grid.Coordinate) [4]grid.Coordinate {
	return [4]grid.Coordinate{c.NewInDir(N), c.NewInDir(E), c.NewInDir(S), c.NewInDir(W)}
}

func countTrails(g grid.Grid, c grid.Coordinate, score *atomic.Int32) {
	cardinals := getCardinals(c)
	next := g.Get(c) + 1

	for _, card := range cardinals {
		if card.InBounds(len(g)-1) && g.Get(card) == next {
			if g.Get(card) == trailHead+end {
				score.Add(1)
				continue
			}
			countTrails(g, card, score)
		}
	}
}
