package day16

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	pattern = regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.+)$`)
)

func part1(input string) (int, error) {
	valves, err := parseValves(input)
	if err != nil {
		return 0, err
	}

	if len(valves) > 64 {
		return 0, errors.New("solution is implemented for not more than 64 valves")
	}

	paths := findAllPaths(valves, 30)

	max := 0
	for _, pressure := range paths {
		if max < pressure {
			max = pressure
		}
	}
	return max, nil
}

func part2(input string) (int, error) {
	valves, err := parseValves(input)
	if err != nil {
		return 0, err
	}

	if len(valves) > 64 {
		return 0, errors.New("solution is implemented for not more than 64 valves")
	}

	paths := findAllPaths(valves, 26)

	// full check takes way too long, it's wise to check the big numbers first
	sorted := make([]BitSet, 0, len(paths))
	for key := range paths {
		sorted = append(sorted, key)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return paths[sorted[i]] > paths[sorted[j]]
	})

	max := 0
	for _, open1 := range sorted {
		pressure1 := paths[open1]
		for _, open2 := range sorted {
			if open1&open2 == 0 {
				if sum := pressure1 + paths[open2]; max < sum {
					max = sum
				} else if sum < max {
					// the rest won't be big enough anyway
					break
				}
			}
		}
	}

	return max, nil
}

type Valve struct {
	label       string
	rate        int
	connections []string
}

type State struct {
	valve    int
	open     BitSet
	pressure int
	time     int
}

type BitSet uint64

func (b BitSet) WithSet(bit int) BitSet {
	return b | (1 << bit)
}

func (b BitSet) IsEmpty(bit int) bool {
	return b&(1<<bit) == 0
}

type Stack []State

func (s *Stack) Push(state State) {
	*s = append(*s, state)
}

func (s *Stack) Pop() State {
	last := len(*s) - 1
	state := (*s)[last]
	*s = (*s)[:last]
	return state
}

// findAllPaths performs DFS, finding all the possible paths within a given time
func findAllPaths(valves []Valve, time int) map[BitSet]int {
	distances := findShortestDistances(valves)

	// Another optimization would be storing just non-zero rates, that would need more complicated starting state calculation
	rates := make([]int, len(valves))
	start := 0
	for i := range valves {
		rates[i] = valves[i].rate
		if valves[i].label == "AA" {
			start = i
		}
	}

	states := make(map[BitSet]int)
	stack := Stack{
		State{
			start,
			0,
			0,
			time,
		},
	}
	for len(stack) != 0 {
		current := stack.Pop()
		if states[current.open] < current.pressure {
			states[current.open] = current.pressure
		}
		if current.time > 0 {
			rate := rates[current.valve]
			if rate != 0 && current.open.IsEmpty(current.valve) {
				time := current.time - 1
				stack.Push(State{
					current.valve,
					current.open.WithSet(current.valve),
					current.pressure + rate*time,
					time,
				})
			} else {
				for next, d := range distances[current.valve] {
					if next != current.valve && rates[next] != 0 && current.open.IsEmpty(next) && current.time > d {
						stack.Push(State{
							next,
							current.open,
							current.pressure,
							current.time - d,
						})
					}
				}
			}
		}
	}
	return states
}

// findShortestDistances finds the minimal distances between each pair of the valves using Floyd-Warshall algorithm.
func findShortestDistances(valves []Valve) [][]int {
	distances := make([][]int, len(valves))
	for i := range distances {
		distances[i] = make([]int, len(valves))
	}
	infinity := len(valves) * len(valves)
	for i := 0; i < len(valves); i++ {
		for j := 0; j < len(valves); j++ {
			if valves[j].label != valves[i].label {
				distances[i][j] = infinity
				for _, next := range valves[j].connections {
					if next == valves[i].label {
						distances[i][j] = 1
						break
					}
				}
			}
		}
	}
	for k := 0; k < len(distances); k++ {
		for i := 0; i < len(distances); i++ {
			for j := 0; j < len(distances); j++ {
				d := distances[i][k] + distances[k][j]
				if d < distances[i][j] {
					distances[i][j] = d
				}
			}
		}
	}
	return distances
}

func parseValves(input string) ([]Valve, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	valves := make([]Valve, len(lines))
	for i, line := range lines {
		submatch := pattern.FindStringSubmatch(line)
		label := submatch[1]
		rate, err := strconv.Atoi(submatch[2])
		if err != nil {
			return nil, err
		}
		connections := strings.Split(submatch[3], ", ")
		valves[i] = Valve{
			label,
			rate,
			connections,
		}
	}
	return valves, nil
}
