package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

const (
	ShapeWidth  = 3
	ShapeLength = 3
)

type Point struct {
	X, Y int
}

type Shape []Point

type Region struct {
	Width       int
	Length      int
	ShapeCounts []int
}

func Part1(r io.Reader) int {
	shapes, regions := parseInput(r)
	numValid := 0
	for _, region := range regions {
		if isRegionValid(shapes, region) {
			numValid++
		}
	}
	return numValid
}

func isRegionValid(shapes []Shape, region Region) bool {
	if !isRegionLargeEnough(shapes, region) {
		return false
	}
	if regionCanFitBoundingBoxes(region) {
		return true
	}
	panic("NP-complete problem")
}

func isRegionLargeEnough(shapes []Shape, region Region) bool {
	totalSpace := 0
	for i, count := range region.ShapeCounts {
		totalSpace += len(shapes[i]) * count
	}
	return totalSpace <= region.Width*region.Length
}

func regionCanFitBoundingBoxes(region Region) bool {
	numBoxes := 0
	for _, count := range region.ShapeCounts {
		numBoxes += count
	}
	capacity := int(region.Width/ShapeWidth) * int(region.Length/ShapeLength)
	return numBoxes <= capacity
}

func parseInput(r io.Reader) ([]Shape, []Region) {
	s := bufio.NewScanner(r)
	var shapes []Shape
	var regions []Region
	re := regexp.MustCompile(`^(\d+):$`)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		if re.MatchString(line) {
			shapes = append(shapes, parseShape(s))
		} else {
			regions = append(regions, parseRegion(line))
		}
	}
	return shapes, regions
}

func parseShape(s *bufio.Scanner) Shape {
	var shape Shape
	y := 0
	width := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		if len(line) > width {
			width = len(line)
		}
		for x, c := range line {
			if c != '#' {
				continue
			}
			shape = append(shape, Point{X: x, Y: y})
		}
		y++
	}
	return shape
}

func parseRegion(s string) Region {
	key, values, found := strings.Cut(s, ":")
	if !found {
		panic("bad input")
	}
	dimensions := strings.Split(key, "x")
	width := must.Get(strconv.Atoi(dimensions[0]))
	length := must.Get(strconv.Atoi(dimensions[1]))
	var shapeCounts []int
	for count := range strings.FieldsSeq(values) {
		shapeCounts = append(shapeCounts, must.Get(strconv.Atoi(count)))
	}
	return Region{
		Width:       width,
		Length:      length,
		ShapeCounts: shapeCounts,
	}
}

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(Part1(strings.NewReader(input)))
}
