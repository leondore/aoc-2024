package grid

import "testing"

func TestMove(t *testing.T) {
	cases := []struct {
		coord Coordinate
		dir   Coordinate
		want  Coordinate
	}{
		{Coordinate{X: 3, Y: 1}, Coordinate{X: 2, Y: -1}, Coordinate{X: 5, Y: 0}},
		{Coordinate{X: 3, Y: 1}, Coordinate{X: 3, Y: 1}, Coordinate{X: 6, Y: 2}},
		{Coordinate{X: 1, Y: 2}, Coordinate{X: 2, Y: 4}, Coordinate{X: 3, Y: 6}},
	}

	for _, c := range cases {
		c.coord.Move(c.dir)

		if c.coord != c.want {
			t.Errorf("got %+v, want %+v", c.coord, c.want)
		}
	}
}

func TestCoordinateInBounds(t *testing.T) {
	cases := []struct {
		coord Coordinate
		limit int
		want  bool
	}{
		{Coordinate{X: 1, Y: 5}, 11, true},
		{Coordinate{X: -1, Y: 11}, 11, false},
		{Coordinate{X: 0, Y: 12}, 11, false},
		{Coordinate{X: 0, Y: 11}, 11, true},
	}

	for _, c := range cases {
		got := c.coord.InBounds(c.limit)

		if got != c.want {
			t.Errorf("Coordinate{%d, %d} in bounds should be %v, got %v", c.coord.X, c.coord.Y, got, c.want)
		}
	}
}
