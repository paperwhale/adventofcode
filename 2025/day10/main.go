package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

type Machine struct {
	Target  int   // int bitmask
	Buttons []int // each button is an int bitmask
}

func Part1(r io.Reader) int {
	machines := parseInput(r)
	presses := 0
	for _, m := range machines {
		presses += solveMachine(m)
	}
	return presses
}

func solveMachine(m Machine) int {
	minCost := len(m.Buttons)
	var solve func(index, state, cost int)
	solve = func(index, state, cost int) {
		if cost >= minCost {
			return
		}
		if index == len(m.Buttons) {
			if state == m.Target {
				minCost = cost
			}
			return
		}
		solve(index+1, state^m.Buttons[index], cost+1)
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
		target := 0
		for i, c := range strings.Trim(tokens[0], "[]") {
			if c == '#' {
				target |= (1 << i)
			}
		}
		buttons := make([]int, 0, len(tokens)-2)
		for _, token := range tokens[1 : len(tokens)-1] {
			button := 0
			for num := range strings.SplitSeq(strings.Trim(token, "()"), ",") {
				button |= (1 << must.Get(strconv.Atoi(num)))
			}
			buttons = append(buttons, button)
		}
		machines = append(machines, Machine{
			Target:  target,
			Buttons: buttons,
		})
	}
	return machines
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input)))
}
