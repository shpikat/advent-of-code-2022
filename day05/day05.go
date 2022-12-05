package day05

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	pattern = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
)

func part1(input string) (string, error) {
	return rearrangeCrates(input, moveOneByOne)
}

func part2(input string) (string, error) {
	return rearrangeCrates(input, moveMultipleAtOnce)
}

func rearrangeCrates(input string, move func(stacks [][]byte, n int, from int, to int)) (string, error) {
	scheme, procedure, found := strings.Cut(input, "\n\n")
	if !found {
		return "", errors.New("expecting two parts in the input")
	}

	lines := strings.Split(scheme, "\n")
	var count int
	last := lines[len(lines)-1]
	for i := len(last) - 1; i >= 0; i-- {
		if last[i] != ' ' {
			// assume the numbers are starting at 1 and increased by 1
			count = int(last[i] - '0')
			break
		}
	}

	stacks := make([][]byte, count)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for i := 0; i < count; i++ {
			index := 1 + i*4
			if index >= len(line) {
				break
			}
			ch := line[index]
			if ch != ' ' {
				stacks[i] = append(stacks[i], ch)
			}
		}
	}

	for _, line := range strings.Split(strings.TrimSpace(procedure), "\n") {
		n, from, to, err := parseStep(line)
		if err != nil {
			return "", err
		}
		// crates numbers are 1-based
		move(stacks, n, from-1, to-1)
	}

	var sb strings.Builder
	for _, s := range stacks {
		err := sb.WriteByte(s[len(s)-1])
		if err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}

func moveOneByOne(stacks [][]byte, n int, from int, to int) {
	length := len(stacks[from])
	for i := 0; i < n; i++ {
		stacks[to] = append(stacks[to], stacks[from][length-1-i])
	}
	stacks[from] = stacks[from][:length-n]
}

func moveMultipleAtOnce(stacks [][]byte, n int, from int, to int) {
	index := len(stacks[from]) - n
	stacks[to] = append(stacks[to], stacks[from][index:]...)
	stacks[from] = stacks[from][:index]
}

func parseStep(line string) (count, from, to int, err error) {
	submatch := pattern.FindStringSubmatch(line)
	count, err = strconv.Atoi(submatch[1])
	if err != nil {
		return
	}
	from, err = strconv.Atoi(submatch[2])
	if err != nil {
		return
	}
	to, err = strconv.Atoi(submatch[3])
	if err != nil {
		return
	}
	return
}
