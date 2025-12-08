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

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input)))
	fmt.Println(Part2(strings.NewReader(input)))
}

func Part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Scan()
	beams := make([]bool, len(s.Text()))
	for i, r := range s.Text() {
		if r == 'S' {
			beams[i] = true
		}
	}
	count := 0
	for s.Scan() {
		for i, r := range s.Text() {
			if !(beams[i] && r == '^') {
				continue
			}
			beams[i] = false
			if i-1 >= 0 {
				beams[i-1] = true
			}
			if i+1 < len(s.Text()) {
				beams[i+1] = true
			}
			count++
		}
	}
	return count
}

func Part2(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Scan()
	beams := make([]int, len(s.Text()))
	for i, r := range s.Text() {
		if r == 'S' {
			beams[i] = 1
		}
	}
	for s.Scan() {
		for i, r := range s.Text() {
			if beams[i] == 0 || r != '^' {
				continue
			}
			if i-1 >= 0 {
				beams[i-1] += beams[i]
			}
			if i+1 < len(s.Text()) {
				beams[i+1] += beams[i]
			}
			beams[i] = 0
		}
	}
	count := 0
	for _, num := range beams {
		count += num
	}
	return count
}
