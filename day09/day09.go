package day09

import (
	"errors"
	"strconv"
	"strings"
)

var (
	motions = map[string]Knot{
		"U": {0, -1},
		"R": {+1, 0},
		"D": {0, +1},
		"L": {-1, 0},
	}
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var head, tail Knot
	tails := make(map[Knot]bool)
	tails[tail] = true
	for _, line := range lines {
		motion, count, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		for i := 0; i < count; i++ {
			head.Move(motion)
			tail.Follow(head)
			tails[tail] = true
		}
	}

	return len(tails), nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	const N = 10
	rope := [N]Knot{}
	tails := make(map[Knot]bool)
	tails[rope[N-1]] = true
	for _, line := range lines {
		motion, count, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		for i := 0; i < count; i++ {
			rope[0].Move(motion)
			for j := 1; j < N; j++ {
				rope[j].Follow(rope[j-1])
			}
			tails[rope[N-1]] = true
		}
	}

	return len(tails), nil
}

type Knot struct {
	x, y int
}

func (c *Knot) Move(motion Knot) {
	c.x += motion.x
	c.y += motion.y
}

func (c *Knot) Follow(head Knot) {
	dx := head.x - c.x
	dy := head.y - c.y

	if Abs(dx) > 1 || Abs(dy) > 1 {
		if dx < 0 {
			c.x--
		} else if dx > 0 {
			c.x++
		}
		if dy < 0 {
			c.y--
		} else if dy > 0 {
			c.y++
		}
	}
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func parseLine(line string) (Knot, int, error) {
	var motion Knot
	first, second, found := strings.Cut(line, " ")
	if !found {
		return motion, 0, errors.New("expecting two values on the line: " + line)
	}
	motion, exists := motions[first]
	if !exists {
		return motion, 0, errors.New("expecting motion U, R, D or L: " + first)
	}
	count, err := strconv.Atoi(second)
	if err != nil {
		return motion, 0, err
	}
	return motion, count, nil
}
