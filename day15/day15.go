package day15

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Reading struct {
	sensor Coordinate
	beacon Coordinate
}

type Coordinate struct {
	x, y int
}

var (
	pattern = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
)

func part1(input string, row int) (int, error) {
	readings, err := parseReadings(input)
	if err != nil {
		return 0, err
	}

	var intervals Intervals
	for _, r := range readings {
		distance := Manhattan(r.sensor, r.beacon)
		dx := distance - Abs(r.sensor.y-row)
		if dx > 0 {
			from := r.sensor.x - dx
			to := r.sensor.x + dx
			if r.beacon.y == row {
				if r.beacon.x == from {
					from++
				} else if r.beacon.x == to {
					to--
				}
			}
			intervals.Add(from, to)
		}
	}

	count := 0
	for _, i := range intervals {
		count += i.to - i.from + 1
	}

	return count, nil
}

func part2(input string, max int) (int, error) {
	readings, err := parseReadings(input)
	if err != nil {
		return 0, err
	}

	// pre-calculate ranges for each sensor
	distances := make(map[Coordinate]int)
	for _, r := range readings {
		distances[r.sensor] = Manhattan(r.sensor, r.beacon)
	}

	type Sensor struct {
		sensor      Coordinate
		distance    int
		touching    int
		overlapping []Coordinate
	}
	sensors := make([]Sensor, len(readings))
	for i, r := range readings {
		distance := distances[r.sensor]
		touching := 0
		var overlapping []Coordinate
		for s, d := range distances {
			if s != r.sensor {
				x := Manhattan(r.sensor, s) - distance - d - 2
				if x < 0 {
					overlapping = append(overlapping, s)
				} else if x == 0 {
					touching++
				}
			}
		}

		sensors[i] = Sensor{
			r.sensor,
			distance,
			touching,
			overlapping,
		}
	}

	// first check where the chances are better
	sort.Slice(sensors, func(i, j int) bool {
		return sensors[i].touching > sensors[j].touching
	})

	edges := []struct {
		start Coordinate
		end   Coordinate
		delta Coordinate
	}{
		{
			Coordinate{-1, 0},
			Coordinate{0, -1},
			Coordinate{+1, -1},
		},
		{
			Coordinate{0, -1},
			Coordinate{+1, 0},
			Coordinate{+1, +1},
		},
		{
			Coordinate{+1, 0},
			Coordinate{0, +1},
			Coordinate{-1, +1},
		},
		{
			Coordinate{0, +1},
			Coordinate{-1, 0},
			Coordinate{-1, -1},
		},
	}
	for _, s := range sensors {
		outside := s.distance + 1
		for _, edge := range edges {
			start := Coordinate{
				s.sensor.x + edge.start.x*outside,
				s.sensor.y + edge.start.y*outside,
			}
			end := Coordinate{
				s.sensor.x + edge.end.x*outside,
				s.sensor.y + edge.end.y*outside,
			}
			c := start
			for c != end {
				if 0 <= c.x && c.x <= max && 0 <= c.y && c.y <= max {
					found := true
					for _, s := range s.overlapping {
						if Manhattan(s, c) <= distances[s] {
							found = false
							break
						}
					}
					if found {
						return c.x*4000000 + c.y, nil
					}
				}
				c.x += edge.delta.x
				c.y += edge.delta.y
			}
		}
	}

	return 0, errors.New("no solution found")
}

func Manhattan(c1, c2 Coordinate) int {
	return Abs(c1.x-c2.x) + Abs(c1.y-c2.y)
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func parseReadings(input string) ([]Reading, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	readings := make([]Reading, len(lines))
	for i, line := range lines {
		submatch := pattern.FindStringSubmatch(line)
		sx, err := strconv.Atoi(submatch[1])
		if err != nil {
			return nil, err
		}
		sy, err := strconv.Atoi(submatch[2])
		if err != nil {
			return nil, err
		}
		bx, err := strconv.Atoi(submatch[3])
		if err != nil {
			return nil, err
		}
		by, err := strconv.Atoi(submatch[4])
		if err != nil {
			return nil, err
		}
		readings[i] = Reading{
			Coordinate{sx, sy},
			Coordinate{bx, by},
		}
	}
	return readings, nil
}
