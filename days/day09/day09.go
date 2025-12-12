package day09

import (
	"fmt"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

type Point struct {
	x, y int
}

type Rect struct {
	minX, maxX, minY, maxY int
}

type Line struct {
	minX, maxX, minY, maxY int
}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day09")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	points := parsePoints(lines)

	if part == 1 {
		maxArea := findLargestRectangle(points)
		fmt.Printf("Part 1: %d\n", maxArea)
	} else {
		maxArea := findLargestValidRectangle(points)
		fmt.Printf("Part 2: %d\n", maxArea)
	}
}

func parsePoints(lines []string) []Point {
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}
	return points
}

func findLargestRectangle(points []Point) int {
	maxArea := 0
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			width := abs(p2.x-p1.x) + 1
			height := abs(p2.y-p1.y) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findLargestValidRectangle(points []Point) int {
	n := len(points)

	lines := make([]Line, 0, n)
	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]

		minX := min(p1.x, p2.x)
		maxX := max(p1.x, p2.x)
		minY := min(p1.y, p2.y)
		maxY := max(p1.y, p2.y)

		lines = append(lines, Line{minX, maxX, minY, maxY})
	}

	maxArea := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			minX := min(p1.x, p2.x)
			maxX := max(p1.x, p2.x)
			minY := min(p1.y, p2.y)
			maxY := max(p1.y, p2.y)

			width := maxX - minX + 1
			height := maxY - minY + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			valid := true
			for _, line := range lines {
				if lineOverlapsRectInterior(line, minX+1, maxX-1, minY+1, maxY-1) {
					valid = false
					break
				}
			}

			if valid {
				maxArea = area
			}
		}
	}

	return maxArea
}

func lineOverlapsRectInterior(line Line, minX, maxX, minY, maxY int) bool {
	if maxX < minX || maxY < minY {
		return false
	}

	lineMaxX := line.maxX
	lineMaxY := line.maxY

	if line.minX > maxX || lineMaxX < minX {
		return false
	}
	if line.minY > maxY || lineMaxY < minY {
		return false
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
