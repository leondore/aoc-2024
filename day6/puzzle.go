package day6

import (
	"strings"

	"github.com/leondore/aoc-2024/utils"
)

const (
	guard    = '^'
	obstacle = '#'
)

type Coordinate struct {
	X, Y int
}

func (c *Coordinate) Move(dir Coordinate) {
	c.X += dir.X
	c.Y += dir.Y
}

func (c *Coordinate) Peek(grid Grid, dir Direction) Coordinate {
	return Coordinate{X: c.X + directions[dir].X, Y: c.Y + directions[dir].Y}
}

func (c *Coordinate) IsAtEdge(grid Grid, dir Direction) bool {
	peek := c.Peek(grid, dir)

	switch dir {
	case N:
		return peek.Y == -1
	case E:
		return peek.X == len(grid)
	case S:
		return peek.Y == len(grid)
	case W:
		return peek.X == -1
	default:
		return false
	}
}

func (c *Coordinate) IsBlocked(grid Grid, dir Direction) bool {
	peek := c.Peek(grid, dir)
	return grid.Get(peek) == obstacle
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

var directions = map[Direction]Coordinate{
	N: {X: 0, Y: -1},
	E: {X: 1, Y: 0},
	S: {X: 0, Y: 1},
	W: {X: -1, Y: 0},
}

type Grid []string

func NewGrid(path string) (Grid, error) {
	return utils.ProcessInput(path)
}

func (g Grid) Get(coord Coordinate) byte {
	return g[coord.Y][coord.X]
}

func (g Grid) FindGuard() Coordinate {
	guardCoords := Coordinate{}

	for idx, row := range g {
		guardX := strings.IndexRune(row, guard)
		if guardX != -1 {
			guardCoords.X, guardCoords.Y = guardX, idx
		}
	}

	return guardCoords
}

func Day6(path string) (int, error) {
	grid, err := NewGrid(path)
	if err != nil {
		return 0, err
	}

	guard := grid.FindGuard()
	dir := N
	visited := map[Coordinate]bool{
		guard: true,
	}

	for {
		guard.Move(directions[dir])
		visited[guard] = true

		if guard.IsAtEdge(grid, dir) {
			break
		}

		if guard.IsBlocked(grid, dir) {
			dir = (dir + 1) % 4
		}
	}

	return len(visited), nil
}
