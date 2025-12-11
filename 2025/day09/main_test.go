package main

import (
	"testing"
)

var example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestPart1(t *testing.T) {
	if got, want := part1(parseInput(example)), 50; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(parseInput(example)), 24; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
