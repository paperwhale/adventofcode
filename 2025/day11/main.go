package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed input.txt
var input string

func Part1(r io.Reader) int {
	conns := parseInput(r)
	const (
		start = "you"
		end   = "out"
	)
	var traverse func(node string) int
	traverse = func(node string) int {
		if node == end {
			return 1
		}
		paths := 0
		for _, neighbor := range conns[node] {
			paths += traverse(neighbor)
		}
		return paths
	}
	return traverse(start)
}

func Part2(r io.Reader) int {
	conns := parseInput(r)
	const (
		start   = "svr"
		target1 = "dac"
		target2 = "fft"
		end     = "out"
	)

	type state struct {
		node  string
		seen1 bool
		seen2 bool
	}
	memo := make(map[state]int)

	var traverse func(node string, seen1, seen2 bool) int
	traverse = func(node string, seen1, seen2 bool) int {
		currState := state{node, seen1, seen2}
		if node == end {
			if !seen1 || !seen2 {
				return 0
			}
			return 1
		}
		if paths, ok := memo[currState]; ok {
			return paths
		}
		paths := 0
		for _, neighbor := range conns[node] {
			nextSeen1 := seen1 || (neighbor == target1)
			nextSeen2 := seen2 || (neighbor == target2)
			paths += traverse(neighbor, nextSeen1, nextSeen2)
		}
		memo[currState] = paths
		return paths
	}
	return traverse(start, false, false)
}

func parseInput(r io.Reader) map[string][]string {
	s := bufio.NewScanner(r)
	m := make(map[string][]string)
	for s.Scan() {
		line := s.Text()
		key, values, found := strings.Cut(line, ":")
		if !found {
			panic("bad input")
		}
		m[key] = strings.Fields(values)
	}
	return m
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input)))
	fmt.Println(Part2(strings.NewReader(input)))
}
