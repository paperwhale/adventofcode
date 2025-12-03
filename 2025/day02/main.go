package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	invalid := 0
	ranges := parseInput(input)
	for _, lr := range ranges {
		l, r := lr[0], lr[1]
		for id := l; id <= r; id++ {
			s := strconv.Itoa(id)
			if s[:len(s)/2] == s[len(s)/2:] {
				invalid += id
			}
		}
	}
	return invalid
}

func Part2(input string) int {
	invalid := 0
	ranges := parseInput(input)
	for _, lr := range ranges {
		l, r := lr[0], lr[1]
		for id := l; id <= r; id++ {
			s := strconv.Itoa(id)
			// s[1:]+s[:len(s)-1] contains every rotation
			// except the original. The original id string
			// has rotational symmetry iff it matches one
			// of the rotations.
			if strings.Contains(s[1:]+s[:len(s)-1], s) {
				invalid += id
			}
		}
	}
	return invalid
}

func parseInput(input string) [][]int {
	var ranges [][]int
	for s := range strings.SplitSeq(input, ",") {
		lr := strings.Split(s, "-")
		l, err := strconv.Atoi(lr[0])
		if err != nil {
			panic("invalid input")
		}
		r, err := strconv.Atoi(lr[1])
		if err != nil {
			panic("invalid input")
		}
		ranges = append(ranges, []int{l, r})
	}
	return ranges
}
