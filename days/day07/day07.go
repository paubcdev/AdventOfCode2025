package day07

import (
	"fmt"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day07")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	if part == 1 {
		splits := countBeamSplits(lines)
		fmt.Printf("Part 1: %d\n", splits)
	} else {
		timelines := countTimelines(lines)
		fmt.Printf("Part 2: %d\n", timelines)
	}
}

func countBeamSplits(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0
	}

	activeBeams := make(map[int]bool)
	activeBeams[startCol] = true

	totalSplits := 0

	for row := 1; row < len(grid); row++ {
		nextBeams := make(map[int]bool)

		for col := range activeBeams {
			if col >= 0 && col < len(grid[row]) && grid[row][col] == '^' {
				totalSplits++
				if col-1 >= 0 {
					nextBeams[col-1] = true
				}
				if col+1 < len(grid[row]) {
					nextBeams[col+1] = true
				}
			} else {
				nextBeams[col] = true
			}
		}

		activeBeams = nextBeams
		if len(activeBeams) == 0 {
			break
		}
	}

	return totalSplits
}

func countTimelines(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	startCol := -1
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startCol = col
			break
		}
	}

	if startCol == -1 {
		return 0
	}

	timelineCounts := make(map[int]int)
	timelineCounts[startCol] = 1

	for row := 1; row < len(grid); row++ {
		nextCounts := make(map[int]int)

		for col, count := range timelineCounts {
			if col >= 0 && col < len(grid[row]) && grid[row][col] == '^' {

				if col-1 >= 0 {
					nextCounts[col-1] += count
				}

				if col+1 < len(grid[row]) {
					nextCounts[col+1] += count
				}
			} else {

				nextCounts[col] += count
			}
		}

		timelineCounts = nextCounts
		if len(timelineCounts) == 0 {
			break
		}
	}

	total := 0
	for _, count := range timelineCounts {
		total += count
	}

	return total
}
