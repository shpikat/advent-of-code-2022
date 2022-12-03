package day03

import (
	"strings"
)

func part1(input string) (int, error) {
	sum := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		compartment := [128]bool{}
		mid := len(line) / 2
		for _, ch := range line[:mid] {
			compartment[ch] = true
		}
		for _, ch := range line[mid:] {
			if compartment[ch] {
				sum += priority(int(ch))
				break
			}
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	sum := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i += 3 {
		compartments := [3][128]bool{}
		for j := 0; j < 3; j++ {
			for _, ch := range lines[i+j] {
				compartments[j][ch] = true
			}
		}
		for ch := 0; ch < 128; ch++ {
			if compartments[0][ch] && compartments[1][ch] && compartments[2][ch] {
				sum += priority(ch)
				break
			}
		}
	}

	return sum, nil
}

func priority(ch int) int {
	var correction int
	if 'a' <= ch && ch <= 'z' {
		correction = 1 - 'a'
	} else {
		correction = 27 - 'A'
	}
	return ch + correction
}
