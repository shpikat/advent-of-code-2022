package day17

import (
	"math/bits"
	"strings"
)

var (
	// the rocks are read bottom to top, meaning index 0 is the bottom line
	rocks = [...]Rock{
		{
			0b0011110,
		},
		{
			0b0001000,
			0b0011100,
			0b0001000,
		},
		{
			0b0011100,
			0b0000100,
			0b0000100,
		},
		{
			0b0010000,
			0b0010000,
			0b0010000,
			0b0010000,
		},
		{
			0b0011000,
			0b0011000,
		},
	}
)

func part1(input string) (int, error) {
	return simulate(strings.TrimSpace(input), 2022), nil
}

func part2(input string) (int, error) {
	return simulate(strings.TrimSpace(input), 1000000000000), nil
}

func simulate(jetsString string, count int) int {
	jets := Jets{
		s: jetsString,
	}

	history := make([][]History, 1<<(bits.Len(uint(len(jetsString)))+bits.Len(uint(len(rocks)))))

	var chamber Chamber
	current := 0
	calculated := 0
	for r := 0; r < count; r++ {
		rock := rocks[current]
		if current == len(rocks)-1 {
			current = 0
		} else {
			current++
		}
		for level := len(chamber) + 3; ; level-- {
			if jets.pushesLeft() {
				rock.moveLeft(chamber, level)
			} else {
				rock.moveRight(chamber, level)
			}
			if chamber.tryLand(rock, level) {
				break
			}
		}

		if calculated == 0 {
			line := chamber[len(chamber)-1]
			key := jets.current<<3 | current
			h := history[key]
			// Allow one full cycle first
			if r > len(rocks)*len(jetsString) && len(h) > 1 {
				for j := len(h) - 1; j >= 0; j-- {
					// Prone to being false positive, in that case the slices of chamber should be compared
					if line == h[j].line {
						cycle := r - h[j].round
						times := (count - r) / cycle
						r += cycle * times
						calculated = (len(chamber) - h[j].height) * times
						break
					}
				}
			}
			history[key] = append(h, History{r, line, len(chamber)})
		}
	}

	return len(chamber) + calculated
}

type Jets struct {
	s       string
	current int
}

func (j *Jets) pushesLeft() bool {
	pushesLeft := j.s[j.current] == '<'
	if j.current == len(j.s)-1 {
		j.current = 0
	} else {
		j.current++
	}
	return pushesLeft
}

type Rock []byte

func (r *Rock) moveLeft(chamber Chamber, level int) {
	next := make(Rock, len(*r))
	for i, b := range *r {
		if b&0b01000000 != 0 {
			return
		}
		next[i] = b << 1
	}
	if !chamber.clashes(next, level) {
		*r = next
	}
}

func (r *Rock) moveRight(chamber Chamber, level int) {
	next := make(Rock, len(*r))
	for i, b := range *r {
		if b&0b00000001 != 0 {
			return
		}
		next[i] = b >> 1
	}
	if !chamber.clashes(next, level) {
		*r = next
	}
}

type Chamber []byte

func (c *Chamber) clashes(rock Rock, level int) bool {
	if level < 0 {
		return true
	}
	for l := 0; l < len(rock) && level+l < len(*c); l++ {
		if (*c)[level+l]&rock[l] != 0 {
			return true
		}
	}
	return false
}

func (c *Chamber) tryLand(rock Rock, level int) bool {
	landed := c.clashes(rock, level-1)
	if landed {
		for l := 0; l < len(rock) && level+l < len(*c); l++ {
			(*c)[level+l] |= rock[l]
		}

		if rest := len(*c) - level; rest < len(rock) {
			*c = append(*c, rock[len(*c)-level:]...)
		}
	}
	return landed
}

type History struct {
	round  int
	line   byte
	height int
}
