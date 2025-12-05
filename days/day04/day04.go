package day04

import (
	"fmt"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day04")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var result int
	if part == 2 {
		result = removeAllAccessibleRolls(lines)
		fmt.Printf("Part 2: %d\n", result)
	} else {
		result = countAccessibleRolls(lines)
		fmt.Printf("Part 1: %d\n", result)
	}
}

func countAccessibleRolls(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			adjacentRolls := 0
			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if grid[nr][nc] == '@' {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				count++
			}
		}
	}

	return count
}

func removeAllAccessibleRolls(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	totalRemoved := 0

	for {
		accessible := findAccessiblePositions(grid)
		if len(accessible) == 0 {
			break
		}

		for _, pos := range accessible {
			grid[pos[0]][pos[1]] = '.'
		}

		totalRemoved += len(accessible)
	}

	return totalRemoved
}

func findAccessiblePositions(grid [][]byte) [][2]int {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var accessible [][2]int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			adjacentRolls := 0
			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if grid[nr][nc] == '@' {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				accessible = append(accessible, [2]int{r, c})
			}
		}
	}

	return accessible
}
