package day18

import (
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

func (c Cube) Move(delta Cube) Cube {
	return Cube{
		c.x + delta.x,
		c.y + delta.y,
		c.z + delta.z,
	}
}

type Box struct {
	bottomLeft, topRight Cube
}

func (b Box) Contains(c Cube) bool {
	return c.x >= b.bottomLeft.x &&
		c.y >= b.bottomLeft.y &&
		c.z >= b.bottomLeft.z &&
		c.x <= b.topRight.x &&
		c.y <= b.topRight.y &&
		c.z <= b.topRight.z
}

var (
	deltas = [...]Cube{
		{-1, 0, 0},
		{+1, 0, 0},
		{0, -1, 0},
		{0, +1, 0},
		{0, 0, -1},
		{0, 0, +1},
	}
)

func part1(input string) (int, error) {
	cubes, err := parseCubes(input)
	if err != nil {
		return 0, err
	}

	index := make(map[Cube]bool, len(cubes))
	for _, c := range cubes {
		index[c] = true
	}

	count := 0
	for _, c := range cubes {
		for _, d := range deltas {
			if !index[c.Move(d)] {
				count++
			}
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	cubes, err := parseCubes(input)
	if err != nil {
		return 0, err
	}

	index := make(map[Cube]bool, len(cubes))
	for _, c := range cubes {
		index[c] = true
	}

	box := getBox(cubes)

	visited := map[Cube]bool{}
	var queue []Cube
	queue = append(queue, box.bottomLeft)
	count := 0
	for len(queue) != 0 {
		batch := queue
		queue = nil
		for _, c := range batch {
			for _, d := range deltas {
				next := c.Move(d)
				if !visited[next] && box.Contains(next) {
					if index[next] {
						count++
					} else {
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
		}
	}

	return count, nil
}

func getBox(cubes []Cube) Box {
	box := Box{cubes[0], cubes[0]}
	for _, c := range cubes[1:] {
		box.bottomLeft.x = min(box.bottomLeft.x, c.x)
		box.bottomLeft.y = min(box.bottomLeft.y, c.y)
		box.bottomLeft.z = min(box.bottomLeft.z, c.z)
		box.topRight.x = max(box.topRight.x, c.x)
		box.topRight.y = max(box.topRight.y, c.y)
		box.topRight.z = max(box.topRight.z, c.z)
	}

	// box surrounds the droplet, must give one extra space
	box.bottomLeft = box.bottomLeft.Move(Cube{-1, -1, -1})
	box.topRight = box.topRight.Move(Cube{1, 1, 1})

	return box
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func parseCubes(input string) ([]Cube, error) {
	var cubes []Cube
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var xyz []int
		for _, s := range strings.Split(line, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			xyz = append(xyz, n)
		}
		cubes = append(cubes, Cube{
			xyz[0],
			xyz[1],
			xyz[2],
		})
	}
	return cubes, nil
}
