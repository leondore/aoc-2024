package day6

import (
	"strings"
	"sync"
	"sync/atomic"

	"github.com/leondore/aoc-2024/utils"
)

const (
	guard           = '^'
	obstacle        = '#'
	specialObstacle = 'O'
	maxRepeatVisits = 5
)

type Coordinate struct {
	X, Y int
}

type Grid []string

func NewGrid(path string) (Grid, error) {
	return utils.ProcessInput(path)
}

func (g Grid) FindGuard() Guard {
	guardCoords := Guard{}

	for idx, row := range g {
		guardX := strings.IndexRune(row, guard)
		if guardX != -1 {
			guardCoords.X, guardCoords.Y = guardX, idx
		}
	}

	return guardCoords
}

type Guard Coordinate

type Scene struct {
	grid  Grid
	guard Guard
	limit int
	dir   Direction
}

func NewScene(path string) (Scene, error) {
	grid, err := NewGrid(path)
	if err != nil {
		return Scene{}, err
	}
	guard := grid.FindGuard()

	return Scene{grid: grid, guard: guard, limit: len(grid), dir: N}, nil
}

func (s *Scene) Get(coord Coordinate) byte {
	return s.grid[coord.Y][coord.X]
}

func (s *Scene) GetGuard() Coordinate {
	return Coordinate(s.guard)
}

func (s *Scene) Move() {
	s.guard.X += directions[s.dir].X
	s.guard.Y += directions[s.dir].Y
}

func (s *Scene) Peek() Coordinate {
	return Coordinate{X: s.guard.X + directions[s.dir].X, Y: s.guard.Y + directions[s.dir].Y}
}

func (s *Scene) IsGuardAtEdge() bool {
	peek := s.Peek()

	switch s.dir {
	case N:
		return peek.Y == -1
	case E:
		return peek.X == s.limit
	case S:
		return peek.Y == s.limit
	case W:
		return peek.X == -1
	default:
		return false
	}
}

func (s *Scene) IsGuardBlocked() bool {
	ahead := s.Get(s.Peek())
	return ahead == obstacle || ahead == specialObstacle
}

func (s *Scene) NewSceneWithObstacle(coord Coordinate) Scene {
	newGrid := make([]string, len(s.grid))
	copy(newGrid, s.grid)

	newGrid[coord.Y] = newGrid[coord.Y][:coord.X] + string(specialObstacle) + newGrid[coord.Y][coord.X+1:]

	return Scene{grid: newGrid, guard: s.guard, limit: s.limit, dir: s.dir}
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

func SimulatePatrolRoute(scene *Scene) (map[Coordinate]int, bool) {
	start := scene.GetGuard()
	visited := map[Coordinate]int{
		start: 1,
	}

	stuck := false

	for {
		scene.Move()
		visited[scene.GetGuard()]++

		if scene.IsGuardAtEdge() {
			break
		}

		for scene.IsGuardBlocked() {
			scene.dir = (scene.dir + 1) % 4
		}

		if visited[scene.GetGuard()] >= maxRepeatVisits {
			stuck = true
			break
		}
	}

	return visited, stuck
}

func Day6(path string) (int, int, error) {
	scene, err := NewScene(path)
	if err != nil {
		return 0, 0, err
	}

	visited, _ := SimulatePatrolRoute(&scene)

	var count atomic.Uint32
	var wg sync.WaitGroup

	resetScene, err := NewScene(path)
	if err != nil {
		return 0, 0, err
	}
	for coord := range visited {
		if resetScene.Get(coord) == guard {
			continue
		}

		wg.Add(1)
		go func() {
			newScene := resetScene.NewSceneWithObstacle(coord)
			_, stuck := SimulatePatrolRoute(&newScene)

			if stuck {
				count.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return len(visited), int(count.Load()), nil
}
