package grid

const EmptyCell = '.'

type Grid []string

func (g Grid) Get(coord Coordinate) byte {
	return g[coord.Y][coord.X]
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) InBounds(limit int) bool {
	return c.X >= 0 && c.X <= limit && c.Y >= 0 && c.Y <= limit
}
