package main

import "testing"

var example = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	if got, want := Part1(example), 13; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := Part2(example), 43; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
