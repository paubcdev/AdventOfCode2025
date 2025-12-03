package day01

import (
	"fmt"
	"strconv"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day01")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	position := 50
	zeroCount := 0

	for _, line := range lines {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing distance: %v\n", err)
			continue
		}

		if part == 2 {
			// To count zeros during rotation
			zeroCount += countZerosDuringRotation(position, direction, distance)
		}

		switch direction {
		case 'L':
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		case 'R':
			position = (position + distance) % 100
		}

		if position == 0 {
			zeroCount++
		}
	}

	if part == 2 {
		fmt.Printf("Part 2: %d\n", zeroCount)
	} else {
		fmt.Printf("Part 1: %d\n", zeroCount)
	}
}

func countZerosDuringRotation(start int, direction byte, distance int) int {
	count := 0

	switch direction {
	case 'L':
		for i := 1; i < distance; i++ {
			pos := (start - i) % 100
			if pos < 0 {
				pos += 100
			}
			if pos == 0 {
				count++
			}
		}
	case 'R':
		for i := 1; i < distance; i++ {
			pos := (start + i) % 100
			if pos == 0 {
				count++
			}
		}
	}

	return count
}
