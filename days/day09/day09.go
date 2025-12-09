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

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day09")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	points := parsePoints(lines)
	maxArea := findLargestRectangle(points)
	fmt.Printf("Part 1: Largest rectangle area = %d\n", maxArea)
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
