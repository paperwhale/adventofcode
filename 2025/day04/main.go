package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var directions = [][]int{
	{0, -1},  // North
	{1, -1},  // North-East
	{1, 0},   // East
	{1, 1},   // South-East
	{0, 1},   // South
	{-1, 1},  // South-West
	{-1, 0},  // West
	{-1, -1}, // North-West
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input string) int {
	count := 0
	grid := parseInput(input)
	for r := range grid {
		for c := range grid[r] {
			if !grid[r][c] {
				continue
			}
			if adjCount(grid, r, c) < 4 {
				count++
			}
		}
	}
	return count
}

func Part2(input string) int {
	count := 0
	grid := parseInput(input)
	for {
		removed := false
		for r := range grid {
			for c := range grid[r] {
				if !grid[r][c] {
					continue
				}
				if adjCount(grid, r, c) < 4 {
					count++
					grid[r][c] = false
					removed = true
				}
			}
		}
		if !removed {
			break
		}
	}
	return count
}

func adjCount(grid [][]bool, r, c int) int {
	count := 0
	for _, dir := range directions {
		newR, newC := r+dir[0], c+dir[1]
		if 0 > newR || 0 > newC || newR >= len(grid) || newC >= len(grid[0]) {
			continue
		}
		if grid[newR][newC] {
			count++
		}
	}
	return count
}

func parseInput(input string) [][]bool {
	var grid [][]bool
	for line := range strings.SplitSeq(input, "\n") {
		row := make([]bool, len(line))
		for i, r := range line {
			if r == '@' {
				row[i] = true
			} else {
				row[i] = false
			}
		}
		grid = append(grid, row)
	}
	return grid
}
