package day8

import (
	"github.com/leondore/aoc-2024/grid"
	"github.com/leondore/aoc-2024/utils"
)

type Map struct {
	Grid      grid.Grid
	Locations map[rune][]grid.Coordinate
}

func (m *Map) FindAntinodes(c1, c2 grid.Coordinate, list map[grid.Coordinate]bool) int {
	limit := len(m.Grid)
	diff := grid.Coordinate{X: c1.X - c2.X, Y: c1.Y - c2.Y}
	count := 0

	for c1.InBounds(limit - 1) {
		list[c1] = true
		c1.Move(diff)
		count++
	}

	for c2.InBounds(limit - 1) {
		list[c2] = true
		c2.Move(grid.Coordinate{X: -diff.X, Y: -diff.Y})
		count++
	}

	return count
}

func (m *Map) CountAntinodes() int {
	antinodes := map[grid.Coordinate]bool{}

	for _, antennas := range m.Locations {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				m.FindAntinodes(antennas[i], antennas[j], antinodes)
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
