package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

// machine represents a single puzzle configuration containing a target light
// state, a set of available buttons, and initial joltage levels.
type machine struct {
	targetLights uint64
	buttons      [][]int
	joltages     []int
}

// buttonCombination represents a pre-calculated effect of a set of buttons.
type buttonCombination struct {
	cost   int
	lights []int
}

func main() {
	machines := parseInput(input)
	fmt.Println(part1(machines))
	fmt.Println(part2(machines))
}

func part1(machines []machine) int {
	totalPresses := 0
	for _, m := range machines {
		totalPresses += minButtonPresses(m)
	}
	return totalPresses
}

func part2(machines []machine) int {
	totalCost := 0
	for _, m := range machines {
		totalCost += minJoltageCost(m)
	}
	return totalCost
}

func minButtonPresses(m machine) int {
	buttonMasks := make([]uint64, len(m.buttons))
	for i, btn := range m.buttons {
		for _, lightIdx := range btn {
			buttonMasks[i] |= 1 << lightIdx
		}
	}

	minCost := len(m.buttons)
	var solve func(index int, currentState uint64, currentCost int)
	solve = func(index int, currentState uint64, currentCost int) {
		if currentCost >= minCost {
			return
		}
		if index == len(m.buttons) {
			if currentState == m.targetLights {
				minCost = currentCost
			}
			return
		}
		solve(index+1, currentState^buttonMasks[index], currentCost+1)
		solve(index+1, currentState, currentCost)
	}

	solve(0, 0, 0)
	return minCost
}

func minJoltageCost(m machine) int {
	combinations := computeButtonCombinations(m)
	memo := make(map[string]int)

	var solve func(currentJoltages []int) int
	solve = func(currentJoltages []int) int {
		if allZero(currentJoltages) {
			return 0
		}

		key := fmt.Sprint(currentJoltages)
		if val, ok := memo[key]; ok {
			return val
		}

		// Calculate the parity mask of the current joltages. We need a button combination
		// that matches this parity to ensure the result after subtraction is divisible by 2.
		targetParity := parityMask(currentJoltages)
		candidates := combinations[targetParity]

		minCost := math.MaxInt
		for _, combo := range candidates {
			nextState, ok := tryReduce(currentJoltages, combo)
			if !ok {
				continue
			}
			res := solve(nextState)
			if res != math.MaxInt {
				total := combo.cost + (2 * res)
				if total < minCost {
					minCost = total
				}
			}
		}

		memo[key] = minCost
		return minCost
	}

	return solve(m.joltages)
}

// computeButtonCombinations generates all possible subsets of buttons and groups
// them by their parity mask.
func computeButtonCombinations(m machine) map[uint64][]buttonCombination {
	combinations := make(map[uint64][]buttonCombination)
	numSubsets := 1 << len(m.buttons)
	numLights := len(m.joltages)

	for mask := range numSubsets {
		combo := makeButtonCombination(mask, m.buttons, numLights)
		pm := parityMask(combo.lights)
		combinations[pm] = append(combinations[pm], combo)
	}
	return combinations
}

func makeButtonCombination(mask int, buttons [][]int, numLights int) buttonCombination {
	lights := make([]int, numLights)
	cost := 0
	for i, btn := range buttons {
		if (mask & (1 << i)) != 0 {
			cost++
			for _, lightIdx := range btn {
				if lightIdx < len(lights) {
					lights[lightIdx]++
				}
			}
		}
	}
	return buttonCombination{cost: cost, lights: lights}
}

func tryReduce(joltages []int, combo buttonCombination) ([]int, bool) {
	next := make([]int, len(joltages))
	for i, v := range joltages {
		diff := v - combo.lights[i]
		if diff < 0 {
			return nil, false
		}
		next[i] = diff / 2
	}
	return next, true
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func parityMask(nums []int) uint64 {
	var mask uint64
	for i, n := range nums {
		if n%2 != 0 {
			mask |= 1 << i
		}
	}
	return mask
}

func parseInput(input string) []machine {
	var machines []machine
	input = strings.TrimSpace(input)
	for line := range strings.SplitSeq(input, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		tokens := strings.Fields(line)
		if len(tokens) < 2 {
			panic(fmt.Sprintf("invalid line format: %s", line))
		}

		// Parse target lights: "[#..#.]"
		lightStr := strings.Trim(tokens[0], "[]")
		var lights uint64
		for i, c := range lightStr {
			if c == '#' {
				lights |= 1 << i
			}
		}

		// Parse buttons: "(1,2)"
		var buttons [][]int
		for _, token := range tokens[1 : len(tokens)-1] {
			buttons = append(buttons, parseCSVInts(strings.Trim(token, "()")))
		}

		// Parse joltages: "{103,34...}"
		joltStr := strings.Trim(tokens[len(tokens)-1], "{}")
		joltages := parseCSVInts(joltStr)

		machines = append(machines, machine{
			targetLights: lights,
			buttons:      buttons,
			joltages:     joltages,
		})
	}

	return machines
}

func parseCSVInts(s string) []int {
	var nums []int
	for s := range strings.SplitSeq(s, ",") {
		nums = append(nums, must.Get(strconv.Atoi(s)))
	}
	return nums
}
