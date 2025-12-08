package main

import (
	"testing"
)

var example = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

func TestPart1(t *testing.T) {
	if got, want := Part1(example), 4277556; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := Part2(example), 3263827; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
