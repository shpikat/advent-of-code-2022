package day11

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	number  int
	items   []int
	change  func(int) int
	divisor int
	yes     int
	no      int
}

func part1(input string) (int, error) {
	monkeys, err := parseMonkeys(input)
	if err != nil {
		return 0, err
	}

	return calculateMonkeyBusiness(monkeys, 20, func(level int) int {
		return level / 3
	})
}

func part2(input string) (int, error) {
	monkeys, err := parseMonkeys(input)
	if err != nil {
		return 0, err
	}

	mod := 1
	for _, m := range monkeys {
		if !isPrime(m.divisor) {
			return 0, fmt.Errorf("divisor is expected to be prime: %d", m.divisor)
		}
		mod *= m.divisor
	}

	return calculateMonkeyBusiness(monkeys, 10000, func(level int) int {
		return level % mod
	})

}

func calculateMonkeyBusiness(monkeys []Monkey, rounds int, adjustLevel func(int) int) (int, error) {
	counts := make([]int, len(monkeys))
	for round := 0; round < rounds; round++ {
		for i, m := range monkeys {
			counts[m.number] += len(m.items)
			for _, level := range m.items {
				next := adjustLevel(m.change(level))
				if next%m.divisor == 0 {
					monkeys[m.yes].items = append(monkeys[m.yes].items, next)
				} else {
					monkeys[m.no].items = append(monkeys[m.no].items, next)
				}
			}
			monkeys[i].items = monkeys[i].items[:0]
		}
	}

	max1, max2 := 0, 0
	for _, c := range counts {
		if c > max1 {
			max1, max2 = c, max1
		} else if c > max2 {
			max2 = c
		}
	}

	return max1 * max2, nil
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func parseMonkeys(input string) ([]Monkey, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	monkeys := make([]Monkey, len(sections))
	for _, section := range sections {
		lines := strings.Split(section, "\n")

		var m Monkey
		var err error
		m.number, err = strconv.Atoi(lines[0][len("Monkey ") : len(lines[0])-1])
		if err != nil {
			return nil, err
		}
		items := strings.Split(lines[1][len("  Starting items: "):], ", ")
		for _, item := range items {
			i, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
			m.items = append(m.items, i)
		}
		operation := lines[2][len("  Operation: new = "):]
		if operation == "old * old" {
			m.change = func(old int) int {
				return old * old
			}
		} else {
			parts := strings.Split(operation, " ")
			v, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, err
			}
			switch parts[1] {
			case "+":
				m.change = func(old int) int {
					return old + v
				}
			case "*":
				m.change = func(old int) int {
					return old * v
				}
			default:
				return nil, errors.New("unsupported operation: " + parts[1])
			}
		}

		m.divisor, err = strconv.Atoi(lines[3][len("  Test: divisible by "):])
		if err != nil {
			return nil, err
		}
		m.yes, err = strconv.Atoi(lines[4][len("    If true: throw to monkey "):])
		if err != nil {
			return nil, err
		}
		m.no, err = strconv.Atoi(lines[5][len("    If false: throw to monkey "):])
		if err != nil {
			return nil, err
		}

		monkeys[m.number] = m
	}
	return monkeys, nil
}
