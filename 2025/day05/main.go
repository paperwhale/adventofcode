package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"slices"
	"strconv"
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
	count := 0
	s := bufio.NewScanner(r)
	fresh := getFresh(s)
	available := getAvailable(s)
	for _, id := range available {
		for _, interval := range fresh {
			if interval[0] <= id && id <= interval[1] {
				count++
				break
			}
		}
	}
	return count
}

func Part2(r io.Reader) int {
	count := 0
	s := bufio.NewScanner(r)
	fresh := getFresh(s)
	slices.SortFunc(fresh, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})
	var merged [][]int
	for _, interval := range fresh {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	for _, interval := range merged {
		count += interval[1] - interval[0] + 1
	}
	return count
}

func getFresh(s *bufio.Scanner) [][]int {
	var fresh [][]int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		strIds := strings.Split(line, "-")
		ids := make([]int, 2)
		for i, strNum := range strIds {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic("bad input")
			}
			ids[i] = num
		}
		fresh = append(fresh, ids)
	}
	return fresh
}

func getAvailable(s *bufio.Scanner) []int {
	var available []int
	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("bad input")
		}
		available = append(available, num)
	}
	return available
}
