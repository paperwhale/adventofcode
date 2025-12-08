package main

import (
	"strings"
	"testing"
)

var example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestPart1(t *testing.T) {
	if got, want := Part1(strings.NewReader(example), 10), 40; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := Part2(strings.NewReader(example)), 25272; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
