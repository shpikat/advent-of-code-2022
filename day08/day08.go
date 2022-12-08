package day08

import (
	"strings"
)

func part1(input string) (int, error) {
	grid := readGrid(input)

	visible := make([][]bool, len(grid))
	// don't check the edges, they are definitely visible
	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]
		visible[i] = make([]bool, len(row))
		threshold := row[0]
		for j := 1; j < len(grid)-1; j++ {
			if row[j] > threshold {
				visible[i][j] = true
				threshold = row[j]

				// no early exits on reaching the maximum tree height are implemented:
				// it is very data-dependent, may incur negative performance impact
				// for introducing yet another conditional
			}
		}
		threshold = row[len(grid)-1]
		for j := len(grid) - 2; j > 0; j-- {
			if row[j] > threshold {
				visible[i][j] = true
				threshold = row[j]
			}
		}
	}
	for i := 1; i < len(grid[0])-1; i++ {
		threshold := grid[0][i]
		for j := 1; j < len(grid)-1; j++ {
			if grid[j][i] > threshold {
				visible[j][i] = true
				threshold = grid[j][i]
			}
		}
		threshold = grid[len(grid)-1][i]
		for j := len(grid) - 2; j > 0; j-- {
			if grid[j][i] > threshold {
				visible[j][i] = true
				threshold = grid[j][i]
			}
		}

	}

	count := (len(grid) + len(grid[0]) - 2) * 2
	for i := 1; i < len(visible)-1; i++ {
		for j := 1; j < len(visible[i])-1; j++ {
			if visible[i][j] {
				count++
			}
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	grid := readGrid(input)

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			current := grid[i][j]
			left, right, up, down := 0, 0, 0, 0
			for j-1-left >= 0 {
				left++
				if grid[i][j-left] >= current {
					break
				}
			}
			for j+1+right < len(grid[i]) {
				right++
				if grid[i][j+right] >= current {
					break
				}
			}
			for i-1-up >= 0 {
				up++
				if grid[i-up][j] >= current {
					break
				}
			}
			for i+1+down < len(grid) {
				down++
				if grid[i+down][j] >= current {
					break
				}
			}

			score := left * right * up * down
			if score > max {
				max = score
			}
		}
	}

	return max, nil
}

func readGrid(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i := range lines {
		line := lines[i]
		grid[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			grid[i][j] = int(line[j] - '0')
		}
	}
	return grid
}
