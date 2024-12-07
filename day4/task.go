package day4

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/leondore/aoc-2024/utils"
)

const word = "XMAS"

type Coordinate [2]int

func (c Coordinate) Move(direction Coordinate) Coordinate {
	return Coordinate{c[0] + direction[0], c[1] + direction[1]}
}

func (c Coordinate) CanMove(direction Coordinate, limit int) bool {
	return c[0]+direction[0] >= 0 && c[0]+direction[0] < limit && c[1]+direction[1] >= 0 && c[1]+direction[1] < limit
}

var directions = [8]Coordinate{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

type Grid []string

func (g Grid) Get(coord Coordinate) byte {
	return g[coord[1]][coord[0]]
}

func Day4(path string) (int, error) {
	gridRaw, err := utils.ProcessInput(path)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %w", err)
	}
	grid := Grid(gridRaw)

	var count atomic.Uint32
	var wg sync.WaitGroup

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			wg.Add(1)
			go func(x, y int) {
				checkAllDirections(&grid, Coordinate{x, y}, &count)
				wg.Done()
			}(x, y)
		}
	}

	wg.Wait()

	return int(count.Load()), nil
}

func checkAllDirections(grid *Grid, coord Coordinate, count *atomic.Uint32) {
	if grid.Get(coord) != word[0] {
		return
	}

	for _, direction := range directions {
		if coord.CanMove(direction, len(*grid)) && checkPossiblePath(grid, coord.Move(direction), direction, 1) {
			count.Add(1)
		}
	}
}

func checkPossiblePath(grid *Grid, coord Coordinate, direction Coordinate, position int) bool {
	if grid.Get(coord) != word[position] {
		return false
	}

	if position == len(word)-1 {
		return true
	}

	if coord.CanMove(direction, len(*grid)) {
		return checkPossiblePath(grid, coord.Move(direction), direction, position+1)
	}

	return false
}
