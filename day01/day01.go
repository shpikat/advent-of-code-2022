package day01

import (
	"strconv"
	"strings"
)

func part1(input string) (int, error) {
	elves := strings.Split(strings.TrimSpace(input), "\n\n")

	max := 0
	for _, elf := range elves {
		current := 0
		for _, food := range strings.Split(elf, "\n") {
			calories, err := strconv.Atoi(food)
			if err != nil {
				return 0, err
			}
			current += calories
		}
		if current > max {
			max = current
		}
	}

	return max, nil
}

func part2(input string) (int, error) {
	elves := strings.Split(strings.TrimSpace(input), "\n\n")

	// Max heap can be used to have O(n*logk) instead of O(n*k),
	// which is quite insignificant in this particular case:
	// k=3 not far from logk ~ 1.5 and n is not big enough to make that
	// a big difference. Code simplicity wins!
	top := [3]int{}
	for _, elf := range elves {
		current := 0
		for _, food := range strings.Split(elf, "\n") {
			calories, err := strconv.Atoi(food)
			if err != nil {
				return 0, err
			}
			current += calories
		}
		if current > top[0] {
			top[0], top[1], top[2] = current, top[0], top[1]
		} else if current > top[1] {
			top[1], top[2] = current, top[1]
		} else if current > top[2] {
			top[2] = current
		}
	}

	total := 0
	for _, v := range top {
		total += v
	}

	return total, nil
}
