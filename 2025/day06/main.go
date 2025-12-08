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

func Part1(s string) int {
	lines := strings.Split(s, "\n")
	ops := strings.Fields(lines[len(lines)-1])
	results := make([]int, len(ops))
	for i := range results {
		if ops[i] == "*" {
			results[i] = 1
		}
	}
	for _, line := range lines[:len(lines)-1] {
		for i, val := range strings.Fields(line) {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic("bad input")
			}
			switch ops[i] {
			case "+":
				results[i] += num
			case "*":
				results[i] *= num
			default:
				panic("unknown operator")
			}
		}
	}
	total := 0
	for _, num := range results {
		total += num
	}
	return total
}

func Part2(s string) int {
	lines := strings.Split(s, "\n")
	ops := strings.Fields(lines[len(lines)-1])
	lines = transpose(lines[:len(lines)-1])
	results := make([]int, len(ops))
	for i := range results {
		if ops[i] == "*" {
			results[i] = 1
		}
	}
	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			i++
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic("bad input")
		}
		switch ops[i] {
		case "+":
			results[i] += num
		case "*":
			results[i] *= num
		default:
			panic("unknown operator")
		}
	}
	total := 0
	for _, num := range results {
		total += num
	}
	return total
}

func transpose(lines []string) []string {
	matrix := make([][]rune, len(lines))
	maxWidth := 0
	for i, line := range lines {
		runes := []rune(line)
		matrix[i] = runes
		if len(runes) > maxWidth {
			maxWidth = len(runes)
		}
	}
	result := make([]string, maxWidth)
	for col := range maxWidth {
		var sb strings.Builder
		for row := range len(matrix) {
			if col < len(matrix[row]) {
				sb.WriteRune(matrix[row][col])
			}
		}
		result[col] = sb.String()
	}
	return result
}
