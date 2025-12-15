package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

type Button []int

type Machine struct {
	Lights   uint // bitmask where each bit corresponds to a light state
	Buttons  []Button
	Joltages []int
}

type ButtonCombination struct {
	Cost     int
	Joltages []int
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input)))
	fmt.Println(Part2(strings.NewReader(input)))
}

func Part1(r io.Reader) int {
	machines := parseInput(r)
	presses := 0
	for _, m := range machines {
		presses += solveMachine(m)
	}
	return presses
}

func Part2(r io.Reader) int {
	machines := parseInput(r)
	presses := 0
	for _, m := range machines {
		presses += configureJoltages(m)
	}
	return presses
}

func configureJoltages(m Machine) int {
	combinations := buttonCombinations(m.Buttons, len(m.Joltages))

	memo := make(map[string]int)
	var solve func([]int) int
	solve = func(joltages []int) int {
		if isZero(joltages) {
			return 0
		}

		key := fmt.Sprint(joltages)
		if val, ok := memo[key]; ok {
			return val
		}

		targetParity := parityMask(joltages)
		candidates := combinations[targetParity]
		minCost := math.MaxInt
		for _, combination := range candidates {
			nextState, ok := tryCombination(joltages, combination)
			if ok {
				res := solve(nextState)
				if res != math.MaxInt {
					total := combination.Cost + (2 * res)
					if total < minCost {
						minCost = total
					}
				}
			}
		}
		memo[key] = minCost
		return minCost
	}

	return solve(m.Joltages)
}

// buttonCombinations generates a map of light masks to button combinations
func buttonCombinations(buttons []Button, numLights int) map[uint][]ButtonCombination {
	combinations := make(map[uint][]ButtonCombination)
	numSubsets := 1 << len(buttons)
	for mask := range numSubsets {
		combination := newButtonCombination(mask, buttons, numLights)
		parityMask := parityMask(combination.Joltages)
		combinations[parityMask] = append(combinations[parityMask], combination)
	}
	return combinations
}

// newButtonCombination creates a new ButtonCombination by selecting a subset of
// buttons using the mask
func newButtonCombination(mask int, buttons []Button, numLights int) ButtonCombination {
	joltages := make([]int, numLights)
	cost := 0
	for i, button := range buttons {
		if (mask & (1 << i)) != 0 {
			cost++
			for _, light := range button {
				joltages[light]++
			}
		}
	}
	return ButtonCombination{Cost: cost, Joltages: joltages}
}

// tryCombination reduces the state of joltages by applying a button combination
func tryCombination(joltages []int, combination ButtonCombination) ([]int, bool) {
	next := make([]int, len(joltages))
	for i, v := range joltages {
		diff := v - combination.Joltages[i]
		if diff < 0 {
			return nil, false
		}
		next[i] = diff / 2
	}
	return next, true
}

// isZero returns whether all elements in nums are zero
func isZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

// parityMask returns a bitmask representing the element-wise parity of nums
func parityMask(nums []int) uint {
	mask := uint(0)
	for i, num := range nums {
		if (num & 1) == 1 {
			mask |= (1 << i)
		}
	}
	return mask
}

func solveMachine(m Machine) int {
	minCost := len(m.Buttons)
	var solve func(index int, state uint, cost int)
	solve = func(index int, state uint, cost int) {
		if cost >= minCost {
			return
		}
		if index == len(m.Buttons) {
			if state == m.Lights {
				minCost = cost
			}
			return
		}
		mask := uint(0)
		for _, light := range m.Buttons[index] {
			mask |= 1 << light
		}
		solve(index+1, state^mask, cost+1)
		solve(index+1, state, cost)
	}
	solve(0, 0, 0)
	return minCost
}

func parseInput(r io.Reader) []Machine {
	var machines []Machine
	s := bufio.NewScanner(r)
	for s.Scan() {
		tokens := strings.Fields(s.Text())
		lights := uint(0)
		for i, c := range strings.Trim(tokens[0], "[]") {
			if c == '#' {
				lights |= (1 << i)
			}
		}
		buttons := make([]Button, 0, len(tokens)-2)
		for _, token := range tokens[1 : len(tokens)-1] {
			var button []int
			for num := range strings.SplitSeq(strings.Trim(token, "()"), ",") {
				button = append(button, must.Get(strconv.Atoi(num)))
			}
			buttons = append(buttons, button)
		}

		var joltages []int
		for num := range strings.SplitSeq(strings.Trim(tokens[len(tokens)-1], "{}"), ",") {
			joltages = append(joltages, must.Get(strconv.Atoi(num)))
		}

		machines = append(machines, Machine{
			Lights:   lights,
			Buttons:  buttons,
			Joltages: joltages,
		})
	}
	return machines
}
