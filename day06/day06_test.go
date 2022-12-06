package day06

import (
	"testing"

	"github.com/shpikat/advent-of-code-2022/internal"
)

const (
	sample1 = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	sample2 = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	sample3 = "nppdvjthqldpwncqszvftbrmjlhg"
	sample4 = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	sample5 = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	part1Sample1 = 7
	part1Sample2 = 5
	part1Sample3 = 6
	part1Sample4 = 10
	part1Sample5 = 11
	part1Answer  = 1816

	part2Sample1 = 19
	part2Sample2 = 23
	part2Sample3 = 23
	part2Sample4 = 29
	part2Sample5 = 26
	part2Answer  = 2625
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample1},
		{"sample 2", sample2, part1Sample2},
		{"sample 3", sample3, part1Sample3},
		{"sample 4", sample4, part1Sample4},
		{"sample 5", sample5, part1Sample5},
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
		want  int
	}{
		{"sample 1", sample1, part2Sample1},
		{"sample 2", sample2, part2Sample2},
		{"sample 3", sample3, part2Sample3},
		{"sample 4", sample4, part2Sample4},
		{"sample 5", sample5, part2Sample5},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input)
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
		fn     func(input string) (int, error)
		answer int
	}{
		{"part1", part1, part1Answer},
		{"part2", part2, part2Answer},
	}

	for _, part := range parts {
		b.Run(part.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := part.fn(input)
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
