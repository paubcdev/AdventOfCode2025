package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

type Shape struct {
	cells [][]bool
}

type Region struct {
	width   int
	height  int
	demands []int
}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day12")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	shapes, regions := parseInput(lines)

	if part == 1 {
		count := 0
		for _, region := range regions {
			if canFit(region, shapes) {
				count++
			}
		}
		fmt.Printf("Part 1: %d\n", count)
	}
}

func parseInput(lines []string) ([]Shape, []Region) {
	shapes := make([]Shape, 0)
	regions := make([]Region, 0)

	i := 0
	shapeIndexRe := regexp.MustCompile(`^(\d+):$`)
	regionRe := regexp.MustCompile(`^(\d+)x(\d+):\s*(.+)$`)

	for i < len(lines) {
		line := lines[i]

		if match := shapeIndexRe.FindStringSubmatch(line); match != nil {
			i++
			shapeLines := make([]string, 0)
			for i < len(lines) && !shapeIndexRe.MatchString(lines[i]) && !regionRe.MatchString(lines[i]) {
				shapeLines = append(shapeLines, lines[i])
				i++
			}

			if len(shapeLines) > 0 {
				cells := make([][]bool, len(shapeLines))
				for r, shapeLine := range shapeLines {
					cells[r] = make([]bool, len(shapeLine))
					for c, ch := range shapeLine {
						cells[r][c] = ch == '#'
					}
				}
				shapes = append(shapes, Shape{cells: cells})
			}
		} else if match := regionRe.FindStringSubmatch(line); match != nil {
			width, _ := strconv.Atoi(match[1])
			height, _ := strconv.Atoi(match[2])
			demandsStr := strings.Fields(match[3])
			demands := make([]int, len(demandsStr))
			for j, d := range demandsStr {
				demands[j], _ = strconv.Atoi(d)
			}
			regions = append(regions, Region{width: width, height: height, demands: demands})
			i++
		} else {
			i++
		}
	}

	return shapes, regions
}

func canFit(region Region, shapes []Shape) bool {
	totalArea := 0
	for shapeIdx, count := range region.demands {
		if shapeIdx < len(shapes) {
			shapeArea := countCells(shapes[shapeIdx])
			totalArea += shapeArea * count
		}
	}

	regionArea := region.width * region.height
	if totalArea > regionArea {
		return false
	}

	grid := make([][]int, region.height)
	for i := range grid {
		grid[i] = make([]int, region.width)
	}

	presents := make([]int, 0)
	for shapeIdx, count := range region.demands {
		for j := 0; j < count; j++ {
			presents = append(presents, shapeIdx)
		}
	}

	return backtrack(grid, presents, shapes, 0)
}

func countCells(shape Shape) int {
	count := 0
	for _, row := range shape.cells {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

func backtrack(grid [][]int, presents []int, shapes []Shape, presentIdx int) bool {
	if presentIdx == len(presents) {
		return true
	}

	shapeIdx := presents[presentIdx]
	presentID := presentIdx + 1

	rotations := generateRotations(shapes[shapeIdx])

	minRow, minCol := findFirstEmpty(grid)
	if minRow == -1 {
		return false
	}

	for row := minRow; row < len(grid); row++ {
		startCol := 0
		if row == minRow {
			startCol = minCol
		}
		for col := startCol; col < len(grid[0]); col++ {
			for _, rotation := range rotations {
				if canPlaceAt(grid, rotation, row, col) {
					place(grid, rotation, row, col, presentID)
					if backtrack(grid, presents, shapes, presentIdx+1) {
						return true
					}
					remove(grid, rotation, row, col)
				}
			}
		}
	}

	return false
}

func findFirstEmpty(grid [][]int) (int, int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				return r, c
			}
		}
	}
	return -1, -1
}

func canPlaceAt(grid [][]int, shape Shape, startRow, startCol int) bool {
	hasAnchor := false
	for r, row := range shape.cells {
		for c, cell := range row {
			if cell {
				gr := startRow + r
				gc := startCol + c
				if gr >= len(grid) || gc >= len(grid[0]) || grid[gr][gc] != 0 {
					return false
				}
				if gr == startRow && gc == startCol {
					hasAnchor = true
				}
			}
		}
	}
	return hasAnchor && grid[startRow][startCol] == 0
}

func generateRotations(shape Shape) []Shape {
	rotations := make([]Shape, 0)
	current := shape

	for i := 0; i < 4; i++ {
		rotations = append(rotations, current)
		rotations = append(rotations, flip(current))
		current = rotate90(current)
	}

	unique := make([]Shape, 0)
	seen := make(map[string]bool)
	for _, r := range rotations {
		key := shapeKey(r)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, r)
		}
	}

	return unique
}

func rotate90(shape Shape) Shape {
	rows := len(shape.cells)
	cols := len(shape.cells[0])
	newCells := make([][]bool, cols)
	for i := range newCells {
		newCells[i] = make([]bool, rows)
		for j := range newCells[i] {
			newCells[i][j] = shape.cells[rows-1-j][i]
		}
	}
	return Shape{cells: newCells}
}

func flip(shape Shape) Shape {
	rows := len(shape.cells)
	cols := len(shape.cells[0])
	newCells := make([][]bool, rows)
	for i := range newCells {
		newCells[i] = make([]bool, cols)
		for j := range newCells[i] {
			newCells[i][j] = shape.cells[i][cols-1-j]
		}
	}
	return Shape{cells: newCells}
}

func shapeKey(shape Shape) string {
	var sb strings.Builder
	for _, row := range shape.cells {
		for _, cell := range row {
			if cell {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func place(grid [][]int, shape Shape, startRow, startCol int, id int) {
	for r, row := range shape.cells {
		for c, cell := range row {
			if cell {
				grid[startRow+r][startCol+c] = id
			}
		}
	}
}

func remove(grid [][]int, shape Shape, startRow, startCol int) {
	for r, row := range shape.cells {
		for c, cell := range row {
			if cell {
				grid[startRow+r][startCol+c] = 0
			}
		}
	}
}
