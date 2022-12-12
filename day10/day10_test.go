package day10

import (
	"strings"
	"testing"

	"github.com/shpikat/advent-of-code-2022/internal"
)

const (
	sample1 = `
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`

	part1Sample = 13140
	part1Answer = 10760

	part2Sample = `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`
	part2Answer = `
####.###...##..###..#..#.####..##..#..#.
#....#..#.#..#.#..#.#..#.#....#..#.#..#.
###..#..#.#....#..#.####.###..#....####.
#....###..#.##.###..#..#.#....#.##.#..#.
#....#....#..#.#....#..#.#....#..#.#..#.
#....#.....###.#....#..#.#.....###.#..#.
`
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part1Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part1(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"sample 1", sample1, part2Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if strings.TrimSpace(got) != strings.TrimSpace(tc.want) {
				t.Errorf("Got:\n%v, want:\n%v", got, tc.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	input := internal.ReadInput(b, "./testdata/input.txt")
	b.Run("part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got, err := part1(input)
			if err != nil {
				b.Errorf("Error: %v", err)
			}
			if got != part1Answer {
				b.Errorf("Got: %v, want: %v", got, part1Answer)
			}
		}
	})
	b.Run("part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got, err := part2(input)
			if err != nil {
				b.Errorf("Error: %v", err)
			}
			if got != part2Answer {
				b.Errorf("Got:\n%v, want:\n%v", got, part2Answer)
			}
		}
	})
}
