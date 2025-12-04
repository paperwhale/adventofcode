package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
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
	output := 0
	banks := parseInput(input)
	for _, bank := range banks {
		i, j1 := maxValue(bank[:len(bank)-1])
		j2 := slices.Max(bank[i+1:])
		output += j1*10 + j2
	}
	return output
}

func Part2(input string) int {
	output := 0
	banks := parseInput(input)
	for _, bank := range banks {
		for i := 11; i >= 0; i-- {
			index, joltage := maxValue(bank[:len(bank)-i])
			bank = bank[index+1:]
			output += joltage * int(math.Pow10(i))
		}
	}
	return output
}

// maxValue returns the index and value of the max value in nums.
func maxValue(nums []int) (int, int) {
	var index, value int
	for i, num := range nums {
		if num > value {
			value = num
			index = i
		}
	}
	return index, value
}

func parseInput(input string) [][]int {
	var lines [][]int
	for s := range strings.SplitSeq(input, "\n") {
		line := make([]int, len(s))
		for i, b := range strings.Split(s, "") {
			num, err := strconv.Atoi(b)
			if err != nil {
				panic("bad input")
			}
			line[i] = num
		}
		lines = append(lines, line)
	}
	return lines
}
