package main

import "testing"

var example = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func TestPart1(t *testing.T) {
	got := part1(example)
	want := 3
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"example", example, 6},
		{"right rotation", []string{"R1000"}, 10},
		{"left rotation", []string{"L1000"}, 10},
		{"right boundary", []string{"R150", "R100"}, 3},
		{"left boundary", []string{"L150", "L100"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
