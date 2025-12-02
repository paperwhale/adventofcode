package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const Start = 50
const Mod = 100

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(input []string) int {
	dial := Start
	count := 0
	for _, s := range input {
		dist, err := parse(s)
		if err != nil {
			panic("parse error")
		}
		// Add Mod to ensure remainder is always positive.
		dial = (((dial + dist) % Mod) + Mod) % Mod
		if dial == 0 {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	dial := Start
	count := 0
	for _, s := range input {
		dist, err := parse(s)
		if err != nil {
			panic("parse error")
		}
		// Case where right rotation crosses zero.
		// If it does not cross, dial + dist < Mod => zeroes = 0.
		zeroes := (dial + dist) / Mod
		// Case where left rotation crosses zero.
		// Same as right case but we flip the sign and add 1 if we cross zero.
		if dial+dist <= 0 {
			zeroes *= -1
			if dial != 0 {
				zeroes++
			}
		}
		dial = (((dial + dist) % Mod) + Mod) % Mod
		count += zeroes
	}
	return count
}

// Parse each command into a clockwise distance.
// Counterclockwise distances are negative.
func parse(s string) (int, error) {
	dist, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, err
	}
	if s[0] == 'L' {
		dist *= -1
	}
	return dist, nil
}
