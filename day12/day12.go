package day12

import (
	"container/heap"
	"errors"
	"strings"
)

const (
	Start = 'S'
	End   = 'E'
)

var (
	deltas = []Coordinate{
		{0, -1},
		{-1, 0},
		{+1, 0},
		{0, +1},
	}
)

func part1(input string) (int, error) {
	grid := strings.Split(strings.TrimSpace(input), "\n")

	var start, end Coordinate
	for y, line := range grid {
		for x, ch := range line {
			if ch == Start {
				start = Coordinate{x, y}
			} else if ch == End {
				end = Coordinate{x, y}
			}
		}
	}

	steps, found := findShortest(grid, []Coordinate{start}, end)
	if found {
		return steps, nil
	}
	return 0, errors.New("no solution found")
}

func part2(input string) (int, error) {
	grid := strings.Split(strings.TrimSpace(input), "\n")

	var starts []Coordinate
	var end Coordinate
	for y, line := range grid {
		for x, ch := range line {
			if ch == Start || ch == 'a' {
				starts = append(starts, Coordinate{x, y})
			} else if ch == End {
				end = Coordinate{x, y}
			}
		}
	}

	steps, found := findShortest(grid, starts, end)
	if found {
		return steps, nil
	}
	return 0, errors.New("no solution found")
}

func findShortest(grid []string, starts []Coordinate, end Coordinate) (int, bool) {
	steps := make(map[Coordinate]int)
	queue := &Heap{}
	for _, start := range starts {
		*queue = append(*queue, State{start, 0})
	}
	heap.Init(queue)
	for len(*queue) > 0 {
		current := heap.Pop(queue).(State)
		if current.position == end {
			return current.steps, true
		}
		best, exists := steps[current.position]
		if !exists || current.steps < best {
			steps[current.position] = current.steps
			elevation := grid[current.position.y][current.position.x]
			threshold := elevation + 1
			if elevation == Start {
				threshold = byte('a')
			}
			for _, delta := range deltas {
				next := current.position.Add(delta)
				if 0 <= next.y && next.y < len(grid) && 0 <= next.x && next.x < len(grid[next.y]) {
					candidate := grid[next.y][next.x]
					if ('a' <= candidate && candidate <= threshold) || (candidate == End && elevation == 'z') {
						heap.Push(queue, State{next, current.steps + 1})
					}
				}
			}
		}
	}
	return 0, false
}

type Coordinate struct {
	x, y int
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{c.x + other.x, c.y + other.y}
}

type State struct {
	position Coordinate
	steps    int
}

type Heap []State

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].steps < h[j].steps }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(state any) {
	*h = append(*h, state.(State))
}

func (h *Heap) Pop() any {
	last := len(*h) - 1
	state := (*h)[last]
	*h = (*h)[:last]
	return state
}
