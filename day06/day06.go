package day06

import (
	"errors"
	"strings"
)

func part1(input string) (int, error) {
	return findMarker(input, 4)
}

func part2(input string) (int, error) {
	return findMarker(input, 14)
}

func findMarker(input string, window int) (int, error) {
	input = strings.TrimSpace(input)

	for i := window; i <= len(input); i++ {
		m := make(map[byte]bool, window)
		for j := i - window; j < i; j++ {
			m[input[j]] = false
		}
		if len(m) == window {
			return i, nil
		}
	}

	return -1, errors.New("no marker found")
}
