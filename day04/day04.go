package day04

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	pattern = regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
)

func part1(input string) (int, error) {
	count := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		a, err := parseAssignment(line)
		if err != nil {
			return 0, err
		}

		containsSecond := a[0] >= a[2] && a[1] <= a[3]
		containsFirst := a[0] <= a[2] && a[1] >= a[3]
		if containsSecond || containsFirst {
			count++
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	count := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		a, err := parseAssignment(line)
		if err != nil {
			return 0, err
		}

		if a[1] >= a[2] && a[3] >= a[0] {
			count++
		}
	}

	return count, nil
}
func parseAssignment(line string) ([]int, error) {
	var assignment []int
	for _, submatch := range pattern.FindStringSubmatch(line)[1:] {
		v, err := strconv.Atoi(submatch)
		if err != nil {
			return nil, err
		}
		assignment = append(assignment, v)
	}
	return assignment, nil
}
