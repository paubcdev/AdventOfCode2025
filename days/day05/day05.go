package day05

import (
	"fmt"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadLines("inputs/day05")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var result int
	if part == 2 {
		result = countAllFreshIDs(lines)
		fmt.Printf("Part 2: %d\n", result)
	} else {
		result = countFreshIngredients(lines)
		fmt.Printf("Part 1: %d\n", result)
	}
}

type ingredientRange struct {
	start int
	end   int
}

func countFreshIngredients(lines []string) int {
	var ranges []ingredientRange
	var ingredients []int
	parsingRanges := true

	for _, line := range lines {
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			// Parse range like "3-5"
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				start, err1 := strconv.Atoi(parts[0])
				end, err2 := strconv.Atoi(parts[1])
				if err1 == nil && err2 == nil {
					ranges = append(ranges, ingredientRange{start: start, end: end})
				}
			}
		} else {
			// Parse ingredient ID
			id, err := strconv.Atoi(line)
			if err == nil {
				ingredients = append(ingredients, id)
			}
		}
	}

	// Count how many ingredients are fresh
	freshCount := 0
	for _, id := range ingredients {
		if isFresh(id, ranges) {
			freshCount++
		}
	}

	return freshCount
}

func isFresh(id int, ranges []ingredientRange) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func countAllFreshIDs(lines []string) int {
	var ranges []ingredientRange

	for _, line := range lines {
		if line == "" {
			break
		}

		// Parse range like "3-5"
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				ranges = append(ranges, ingredientRange{start: start, end: end})
			}
		}
	}

	// Merge overlapping ranges and count total IDs
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start position
	// Simple bubble sort since we don't have many ranges
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].start < ranges[i].start {
				ranges[i], ranges[j] = ranges[j], ranges[i]
			}
		}
	}

	// Merge overlapping ranges and count
	totalIDs := 0
	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= currentEnd+1 {
			// Overlapping or adjacent, merge
			if ranges[i].end > currentEnd {
				currentEnd = ranges[i].end
			}
		} else {
			// Non-overlapping, count current range and start new one
			totalIDs += currentEnd - currentStart + 1
			currentStart = ranges[i].start
			currentEnd = ranges[i].end
		}
	}

	// Count the last range
	totalIDs += currentEnd - currentStart + 1

	return totalIDs
}
