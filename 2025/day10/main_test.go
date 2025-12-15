package main

import (
	"testing"
)

var example = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

func TestPart1(t *testing.T) {
	machines := parseInput(example)
	if got, want := part1(machines), 7; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	machines := parseInput(example)
	if got, want := part2(machines), 33; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
