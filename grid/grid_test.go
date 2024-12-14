package grid

import "testing"

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
