package main

import "testing"

var example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	if got, want := Part1(example), 357; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := Part2(example), 3121910778619; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
