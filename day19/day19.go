package day19

import (
	"math"
	"regexp"
	"strconv"
)

const (
	Ore int = iota
	Clay
	Obsidian
	Geode

	MaterialsCount = 4

	// ValueWidth is chosen to be wide enough to fit the values for the given task, and also not letting the full data type
	// exceed 32 bits. This way we can use uint reluctantly instead of selecting uint32 or uint64 explicitly.
	ValueWidth = 8
	ValueMask  = (1 << ValueWidth) - 1
)

var (
	pattern = regexp.MustCompile(`(?m)^Blueprint (\d+):\s+Each ore robot costs (\d+) ore\.\s+Each clay robot costs (\d+) ore\.\s+Each obsidian robot costs (\d+) ore and (\d+) clay\.\s+Each geode robot costs (\d+) ore and (\d+) obsidian\.$`)
)

func part1(input string) (int, error) {
	blueprints, err := parseBlueprints(input)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, blueprint := range blueprints {
		result += blueprint.id * blueprint.FindMaxGeodes(24)
	}
	return result, nil
}

func part2(input string) (int, error) {
	blueprints, err := parseBlueprints(input)
	if err != nil {
		return 0, err
	}

	count := 3
	if len(blueprints) < 3 {
		count = len(blueprints)
	}

	result := 1
	for _, blueprint := range blueprints[:count] {
		result *= blueprint.FindMaxGeodes(32)
	}
	return result, nil
}

// Materials stores the quantity for all the given materials.
//
// Storing the data compressed in uint provides ~30-40% speedup over storing as array of ints.
type Materials uint

func NewMaterials(quantity int, material int) Materials {
	return Materials(quantity << (material * ValueWidth))
}

func (m Materials) Get(index int) int {
	return int((m >> (index * ValueWidth)) & ValueMask)
}

func (m Materials) IsEnough(bill Materials) bool {
	for i := 0; i < MaterialsCount; i++ {
		if (m & ValueMask) < (bill & ValueMask) {
			return false
		}
		m >>= ValueWidth
		bill >>= ValueWidth
	}
	return true
}

type Blueprint struct {
	id     int
	robots [MaterialsCount]Materials
}

func (b Blueprint) FindMaxGeodes(timePeriod int) int {
	max := 0

	// As the robot construction is not instant and takes one minute, there is no need to have more robots of each type
	// than the maximum required to construct any other robot.
	var thresholds [MaterialsCount]int
	for _, bill := range b.robots {
		for material := 0; material < MaterialsCount; material++ {
			quantity := bill.Get(material)
			if quantity > thresholds[material] {
				thresholds[material] = quantity
			}
		}
	}
	thresholds[Geode] = math.MaxInt

	stack := Stack{
		State{
			robots:    NewMaterials(1, Ore),
			countdown: timePeriod,
		}}
	// There is no need to track visited states (major performance hit), as we advance on known paths only
	for len(stack) != 0 {
		current := stack.Pop()
		if current.countdown == 0 {
			geodes := current.inventory.Get(Geode)
			if geodes > max {
				max = geodes
			}
		} else {
			// Prefer "calculate-when-next-robot-is-built" to "what-robot-to-build-at-this-step" to create fewer branches
			for r, bill := range b.robots {
				if current.robots.Get(r) < thresholds[r] {
					next := current

					canBuild := false
					for !canBuild && next.countdown > 0 && next.EstimateMaximum() > max {
						next.countdown--
						canBuild = next.inventory.IsEnough(bill)
						next.inventory += next.robots
					}

					if canBuild {
						next.robots += NewMaterials(1, r)
						next.inventory -= bill
						stack.Push(next)
					} else if next.countdown == 0 {
						// No new robot is built, but we need to account the resources we collected on the way
						stack.Push(next)
					}
				}
			}
		}
	}

	return max
}

type State struct {
	robots    Materials
	inventory Materials
	countdown int
}

func (s State) EstimateMaximum() int {
	current := s.inventory.Get(Geode)
	planned := s.robots.Get(Geode) * s.countdown
	possible := ((s.countdown - 1) * s.countdown) / 2
	return current + planned + possible
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

func parseBlueprints(input string) ([]Blueprint, error) {
	var blueprints []Blueprint
	for _, submatch := range pattern.FindAllStringSubmatch(input, -1) {
		var b Blueprint
		var err error
		b.id, err = strconv.Atoi(submatch[1])
		if err != nil {
			return nil, err
		}
		oreOre, err := strconv.Atoi(submatch[2])
		if err != nil {
			return nil, err
		}
		clayOre, err := strconv.Atoi(submatch[3])
		if err != nil {
			return nil, err
		}
		obsidianOre, err := strconv.Atoi(submatch[4])
		if err != nil {
			return nil, err
		}
		obsidianClay, err := strconv.Atoi(submatch[5])
		if err != nil {
			return nil, err
		}
		geodeOre, err := strconv.Atoi(submatch[6])
		if err != nil {
			return nil, err
		}
		geodeObsidian, err := strconv.Atoi(submatch[7])
		if err != nil {
			return nil, err
		}

		b.robots[Ore] = NewMaterials(oreOre, Ore)
		b.robots[Clay] = NewMaterials(clayOre, Ore)
		b.robots[Obsidian] = NewMaterials(obsidianOre, Ore) | NewMaterials(obsidianClay, Clay)
		b.robots[Geode] = NewMaterials(geodeOre, Ore) | NewMaterials(geodeObsidian, Obsidian)

		blueprints = append(blueprints, b)
	}
	return blueprints, nil
}
