package helpers

import "strings"

type Coord struct {
	X int
	Y int
}

func (c *Coord) RotateCW() {
	if c.X == 1 {
		c.X = 0
		c.Y = -1
	} else if c.Y == -1 {
		c.X = -1
		c.Y = 0
	} else if c.X == -1 {
		c.X = 0
		c.Y = 1
	} else if c.Y == 1 {
		c.X = 1
		c.Y = 0
	}
}

func (c *Coord) RotateCCW() {
	if c.X == 1 {
		c.X = 0
		c.Y = 1
	} else if c.Y == 1 {
		c.X = -1
		c.Y = 0
	} else if c.X == -1 {
		c.X = 0
		c.Y = -1
	} else if c.Y == -1 {
		c.X = 1
		c.Y = 0
	}
}

type Grid [][]rune

func (g *Grid) Init(input string) bool {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		var runes []rune
		for _, r := range line {
			runes = append(runes, r)
		}
		(*g) = append((*g), runes)
	}

	return true
}

func (g Grid) GetPointDir(x int, y int, dx int, dy int) (rune, bool) {
	return g.GetPointCoord(x, y, Coord{dx, dy})
}

func (g Grid) GetPointCoord(x int, y int, c Coord) (rune, bool) {
	x += c.X
	y += c.Y
	return g.GetPoint(x, y)
}

func (g Grid) GetPoint(x int, y int) (rune, bool) {
	if x < 0 || y < 0 || x == len(g) || y == len(g[0]) {
		return ' ', false
	}

	return g[x][y], true
}

func (g Grid) DeepCopy() Grid {
	newGrid := make(Grid, len(g))
	for i := range g {
		newGrid[i] = make([]rune, len(g[i]))
		copy(newGrid[i], g[i])
	}
	return newGrid
}
