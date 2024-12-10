package day4

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/leondore/aoc-2024/utils"
)

const word = "XMAS"

const (
	X int = iota
	M
	A
	S
)

type Coordinate [2]int

func (c Coordinate) Move(direction Coordinate) Coordinate {
	return Coordinate{c[0] + direction[0], c[1] + direction[1]}
}

func (c Coordinate) GetXCoords() (Coordinate, Coordinate, Coordinate, Coordinate) {
	return c.Move(directions[NE]), c.Move(directions[SE]), c.Move(directions[SW]), c.Move(directions[NW])
}

// Deprecated: Used in part 1 - no longer required
func (c Coordinate) CanMove(direction Coordinate, limit int) bool {
	return c[0]+direction[0] >= 0 && c[0]+direction[0] < limit && c[1]+direction[1] >= 0 && c[1]+direction[1] < limit
}

type Direction int

const (
	NE Direction = iota
	SE
	SW
	NW
)

var directions = map[Direction]Coordinate{
	NE: {1, -1},
	SE: {1, 1},
	SW: {-1, 1},
	NW: {-1, -1},
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

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid)-1; x++ {
			if grid.Get(Coordinate{x, y}) != word[A] {
				continue
			}

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
	ne, se, sw, nw := coord.GetXCoords()
	if isSideValid(grid, ne, sw) && isSideValid(grid, nw, se) {
		count.Add(1)
	}
}

func isSideValid(grid *Grid, start Coordinate, end Coordinate) bool {
	startChar := grid.Get(start)
	endChar := grid.Get(end)
	return (startChar == word[M] && endChar == word[S]) || (startChar == word[S] && endChar == word[M])
}

// Deprecated: Used in part 1 - no longer required
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
