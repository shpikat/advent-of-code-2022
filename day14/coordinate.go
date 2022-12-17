package day14

import (
	"fmt"

	"github.com/shpikat/advent-of-code-2022/utils"
)

type Coordinate struct {
	x, y int
}

func (c Coordinate) Move(delta Coordinate) Coordinate {
	return Coordinate{
		c.x + delta.x,
		c.y + delta.y,
	}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

var void utils.Void

type Coordinates map[Coordinate]utils.Void

func (s Coordinates) Add(c Coordinate) {
	s[c] = void
}

func (s Coordinates) Remove(c Coordinate) {
	delete(s, c)
}

func (s Coordinates) Has(c Coordinate) bool {
	_, exists := s[c]
	return exists
}
