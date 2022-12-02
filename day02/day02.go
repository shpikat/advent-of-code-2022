package day02

import (
	"fmt"
	"strings"
)

type Shape int

func (s Shape) score() int {
	return int(s)
}

const (
	None     Shape = 0
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

var (
	firstColumnCodes = map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	beaters = [4]Shape{None, Paper, Scissors, Rock}
	losers  = [4]Shape{None, Scissors, Rock, Paper}
)

func part1(input string) (int, error) {
	secondColumnCodes := map[string]Shape{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	score := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		first, second, ok := strings.Cut(line, " ")
		if !ok {
			return 0, fmt.Errorf("unexpected format: %s", line)
		}
		opponent, ok := firstColumnCodes[first]
		if !ok {
			return 0, fmt.Errorf("expecting only A, B or C, got: %s", first)
		}
		you, ok := secondColumnCodes[second]
		if !ok {
			return 0, fmt.Errorf("expecting only X, Y or Z, got: %s", second)
		}
		score += you.score()
		if opponent == you {
			score += 3
		} else if beaters[opponent] == you {
			score += 6
		}
	}

	return score, nil
}

func part2(input string) (int, error) {
	score := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		first, second, ok := strings.Cut(line, " ")
		if !ok {
			return 0, fmt.Errorf("unexpected format: %s", line)
		}
		opponent, ok := firstColumnCodes[first]
		if !ok {
			return 0, fmt.Errorf("expecting only A, B or C, got: %s", first)
		}
		var you Shape
		switch second {
		case "X":
			you = losers[opponent]
		case "Y":
			you = opponent
			score += 3
		case "Z":
			you = beaters[opponent]
			score += 6
		default:
			return 0, fmt.Errorf("expecting only X, Y or Z, got: %s", second)
		}
		score += you.score()
	}

	return score, nil
}
