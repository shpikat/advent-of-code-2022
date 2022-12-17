package day14

import (
	"errors"
	"strconv"
	"strings"
)

var (
	source = Coordinate{500, 0}

	deltas = []Coordinate{
		{0, +1},
		{-1, +1},
		{+1, +1},
	}
)

func part1(input string) (int, error) {
	rocks, err := parseRocks(input)
	if err != nil {
		return 0, err
	}

	bottom := source.y
	for c := range rocks {
		if c.y > bottom {
			bottom = c.y
		}
	}

	sand := make(Coordinates)
	for {
		current := source

		for {
			var next Coordinate
			found := false
			for _, delta := range deltas {
				next = current.Move(delta)
				if !rocks.Has(next) && !sand.Has(next) {
					found = true
					break
				}
			}

			if !found {
				sand.Add(current)
				break
			}

			if next.y == bottom {
				return len(sand), nil
			}

			current = next
		}
	}
}

func part2(input string) (int, error) {
	rocks, err := parseRocks(input)
	if err != nil {
		return 0, err
	}

	bottom := source.y
	for c := range rocks {
		if c.y > bottom {
			bottom = c.y
		}
	}
	bottom += 2

	// The sand won't go farther than that
	for x := source.x - bottom; x <= source.x+bottom; x++ {
		rocks.Add(Coordinate{x, bottom})
	}

	sand := make(Coordinates)
	for {
		current := source

		for {
			var next Coordinate
			found := false
			for _, delta := range deltas {
				next = current.Move(delta)
				if !rocks.Has(next) && !sand.Has(next) {
					found = true
					break
				}
			}

			if !found {
				if current == source {
					// add 1 instead of adding the source to the rest of the sand
					return len(sand) + 1, nil
				}
				sand.Add(current)
				break
			}

			current = next
		}
	}
}

func parseRocks(input string) (Coordinates, error) {
	rocks := make(Coordinates)
	for _, rock := range strings.Split(strings.TrimSpace(input), "\n") {
		points := strings.Split(rock, " -> ")
		previous, err := parseCoordinate(points[0])
		if err != nil {
			return nil, err
		}
		for _, point := range points[1:] {
			current, err := parseCoordinate(point)
			if err != nil {
				return nil, err
			}
			if previous.x == current.x {
				from, to := previous.y, current.y
				if previous.y > current.y {
					from, to = to, from
				}
				for y := from; y <= to; y++ {
					rocks.Add(Coordinate{current.x, y})
				}
			} else {
				from, to := previous.x, current.x
				if previous.x > current.x {
					from, to = to, from
				}
				for x := from; x <= to; x++ {
					rocks.Add(Coordinate{x, current.y})
				}
			}
			previous = current
		}
	}

	return rocks, nil
}

func parseCoordinate(point string) (Coordinate, error) {
	var c Coordinate
	x, y, found := strings.Cut(point, ",")
	if !found {
		return c, errors.New("unexpected coordinate format: " + point)
	}
	var err error
	c.x, err = strconv.Atoi(x)
	if err != nil {
		return c, err
	}
	c.y, err = strconv.Atoi(y)
	if err != nil {
		return c, err
	}
	return c, nil
}
