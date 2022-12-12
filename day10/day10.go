package day10

import (
	"errors"
	"strconv"
	"strings"
)

type Command struct {
	fn     func(int, int) int
	cycles int
}

var (
	commands = map[string]Command{
		"noop": {
			func(x int, v int) int {
				return x
			},
			1,
		},
		"addx": {
			func(x int, v int) int {
				return x + v
			},
			2,
		},
	}
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	x := 1
	cycle := 0
	stopCycle := 20
	sum := 0
	for _, line := range lines {
		command, v, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		next := cycle + command.cycles
		if next >= stopCycle {
			sum += stopCycle * x
			stopCycle += 40
		}
		cycle = next
		x = command.fn(x, v)
	}

	return sum, nil
}

func part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	const (
		H = 6
		W = 40
	)
	crt := [H * W]byte{}
	x := 1
	cycle := 0
	for _, line := range lines {
		command, v, err := parseLine(line)
		if err != nil {
			return "", err
		}

		end := cycle + command.cycles
		for cycle < end {
			pixel := cycle % W
			if x-1 <= pixel && pixel <= x+1 {
				crt[cycle] = '#'
			} else {
				crt[cycle] = '.'
			}
			cycle++
		}
		x = command.fn(x, v)
	}

	var sb strings.Builder
	for start := 0; start < H*W; start += W {
		sb.WriteString(string(crt[start : start+W]))
		sb.WriteByte('\n')
	}
	return sb.String(), nil
}

func parseLine(line string) (Command, int, error) {
	var command Command
	first, second, found := strings.Cut(line, " ")
	command, exists := commands[first]
	if !exists {
		return command, 0, errors.New("unexpected command: " + first)
	}
	var v int
	if found {
		var err error
		v, err = strconv.Atoi(second)
		if err != nil {
			return command, 0, err
		}
	}
	return command, v, nil
}
