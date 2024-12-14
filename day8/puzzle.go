package day8

import (
	"github.com/leondore/aoc-2024/grid"
	"github.com/leondore/aoc-2024/utils"
)

type Map struct {
	Grid      grid.Grid
	Locations map[rune][]grid.Coordinate
}

func (m *Map) FindAntinodes(c1, c2 grid.Coordinate) (grid.Coordinate, grid.Coordinate) {
	limit := len(m.Grid)
	diff := grid.Coordinate{X: c1.X - c2.X, Y: c1.Y - c2.Y}

	an1 := grid.Coordinate{X: c1.X + diff.X, Y: c1.Y + diff.Y}
	if !an1.InBounds(limit - 1) {
		an1 = grid.Coordinate{}
	}
	an2 := grid.Coordinate{X: c2.X - diff.X, Y: c2.Y - diff.Y}
	if !an2.InBounds(limit - 1) {
		an2 = grid.Coordinate{}
	}

	return an1, an2
}

func (m *Map) CountAntinodes() int {
	antinodes := map[grid.Coordinate]bool{}

	for _, antennas := range m.Locations {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				an1, an2 := m.FindAntinodes(antennas[i], antennas[j])
				if an1 != (grid.Coordinate{}) {
					antinodes[an1] = true
				}
				if an2 != (grid.Coordinate{}) {
					antinodes[an2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func NewMap(g grid.Grid) Map {
	locations := map[rune][]grid.Coordinate{}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			cell := rune(g.Get(grid.Coordinate{X: x, Y: y}))
			if cell != grid.EmptyCell {
				if _, ok := locations[cell]; !ok {
					locations[cell] = []grid.Coordinate{}
				}
				locations[cell] = append(locations[cell], grid.Coordinate{X: x, Y: y})
			}
		}
	}

	return Map{Grid: g, Locations: locations}
}

func Day8(path string) (int, error) {
	input, err := utils.ProcessInput(path)
	if err != nil {
		return 0, err
	}

	m := NewMap(grid.Grid(input))
	return m.CountAntinodes(), nil
}
