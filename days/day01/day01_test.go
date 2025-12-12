package day01

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay01(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{part: 1, expected: 3},
		{part: 2, expected: 6},
	}

	for _, tt := range tests {
		t.Run("Part"+strconv.Itoa(tt.part), func(t *testing.T) {
			lines, err := utils.ReadNonEmptyLines("../test_inputs/testday01")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			result := calculateResult(lines, tt.part)

			if result != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, result)
			}
		})
	}
}

func calculateResult(lines []string, part int) int {
	position := 50
	zeroCount := 0

	for _, line := range lines {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if part == 2 {
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

	return zeroCount
}
