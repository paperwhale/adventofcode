package main

import (
	"strings"
	"testing"
)

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	r := strings.NewReader(example)
	if got, want := Part1(r), 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(example)
	if got, want := Part2(r), 14; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
