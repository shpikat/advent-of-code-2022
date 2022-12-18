package day15

import (
	"testing"

	"github.com/shpikat/advent-of-code-2022/internal"
)

const (
	sample1 = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

	part1Sample = 26
	part1Answer = 4861076

	part2Sample = 56000011
	part2Answer = 10649103160102
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		row   int
		want  int
	}{
		{"sample 1", sample1, 10, part1Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), 2000000, part1Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part1(tc.input, tc.row)
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
		max   int
		want  int
	}{
		{"sample 1", sample1, 20, part2Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), 4000000, part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input, tc.max)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	input := internal.ReadInput(b, "./testdata/input.txt")
	parts := []struct {
		name   string
		fn     func(input string, arg int) (int, error)
		arg    int
		answer int
	}{
		{"part1", part1, 2000000, part1Answer},
		{"part2", part2, 4000000, part2Answer},
	}

	for _, part := range parts {
		b.Run(part.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := part.fn(input, part.arg)
				if err != nil {
					b.Errorf("Error: %v", err)
				}
				if got != part.answer {
					b.Errorf("Got: %v, want: %v", got, part.answer)
				}
			}
		})
	}
}
