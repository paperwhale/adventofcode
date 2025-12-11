package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/paperwhale/adventofcode/internal/must"
)

//go:embed input.txt
var input string

// point represents a point in 2D space.
type point struct {
	x, y int
}

// segment represents a line segment between two points.
type segment struct{ p1, p2 point }

// rayCast returns true if a ray cast from point p intersects the segment s.
func (s segment) rayCast(p point) bool {
	return ((s.p1.y > p.y) != (s.p2.y > p.y)) &&
		(p.x < (s.p2.x-s.p1.x)*(p.y-s.p1.y)/(s.p2.y-s.p1.y)+s.p1.x)
}

// rectangle is defined by its bottom left and upper right corners.
type rectangle struct {
	min, max point
}

// newRectangle creates a new rectangle from any two opposite corners.
func newRectangle(p1, p2 point) rectangle {
	return rectangle{
		min: point{min(p1.x, p2.x), min(p1.y, p2.y)},
		max: point{max(p1.x, p2.x), max(p1.y, p2.y)},
	}
}

// area returns the area of the rectangle.
func (r rectangle) area() int {
	return (r.max.x - r.min.x + 1) * (r.max.y - r.min.y + 1)
}

// contains returns true if the point p is inside the rectangle r.
func (r rectangle) contains(p point) bool {
	return r.min.x < p.x && p.x < r.max.x && r.min.y < p.y && p.y < r.max.y
}

func (r rectangle) center() point {
	return point{
		x: (r.min.x + r.max.x) / 2,
		y: (r.min.y + r.max.y) / 2,
	}
}

func main() {
	points := parseInput(input)
	fmt.Println(part1(points))
	fmt.Println(part2(points))
}

func part1(points []point) int {
	var maxArea int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			r := newRectangle(p1, p2)
			if a := r.area(); a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
}

func part2(points []point) int {
	var maxArea int
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			r := newRectangle(p1, p2)
			if !isValid(r, points) {
				continue
			}
			if a := r.area(); a > maxArea {
				maxArea = a
			}
		}
	}
	return maxArea
}

func isValid(r rectangle, points []point) bool {
	// Candidate rectangle must not contain any other points.
	if slices.ContainsFunc(points, r.contains) {
		return false
	}
	segments := toSegments(points)
	// Check that any point within the rectangle is inside the polygon.
	if !isPointInPolygon(r.center(), segments) {
		return false
	}
	// Ensure no edges of polygon intersect the rectangle edges.
	if intersectsRectangle(r, segments) {
		return false
	}
	return true
}

// toSegments converts a slice of points into a slice of segments connecting them.
func toSegments(points []point) []segment {
	segments := make([]segment, len(points))
	for i := range points {
		segments[i] = segment{points[i], points[(i+1)%len(points)]}
	}
	return segments
}

// isPointInPolygon returns true if the point is contained within the polygon.
func isPointInPolygon(p point, polygon []segment) bool {
	inside := false
	for _, s := range polygon {
		if s.rayCast(p) {
			inside = !inside
		}
	}
	return inside
}

// intersectsRectangle returns true if any segment intersects the edges
// of the rectangle r.
func intersectsRectangle(r rectangle, polygon []segment) bool {
	for _, s := range polygon {
		// Vertical slice check
		if s.p1.x == s.p2.x && s.p1.x > r.min.x && s.p1.x < r.max.x &&
			min(s.p1.y, s.p2.y) <= r.min.y && max(s.p1.y, s.p2.y) >= r.max.y {
			return true
		}
		// Horizontal slice check
		if s.p1.y == s.p2.y && s.p1.y > r.min.y && s.p1.y < r.max.y &&
			min(s.p1.x, s.p2.x) <= r.min.x && max(s.p1.x, s.p2.x) >= r.max.x {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(input string) []point {
	var points []point
	for line := range strings.Lines(strings.TrimSpace(input)) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		points = append(points, point{
			x: must.Get(strconv.Atoi(nums[0])),
			y: must.Get(strconv.Atoi(nums[1])),
		})
	}
	return points
}
