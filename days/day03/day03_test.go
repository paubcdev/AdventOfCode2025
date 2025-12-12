package day03

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay03(t *testing.T) {
	tests := []struct {
		part     int
		expected int64
	}{
		{part: 1, expected: 357},
		{part: 2, expected: 3121910778619},
	}

	for _, tt := range tests {
		t.Run("Part"+strconv.Itoa(tt.part), func(t *testing.T) {
			lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday03")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			totalJoltage := int64(0)

			for _, line := range lines {
				var maxJoltage int64
				if tt.part == 2 {
					maxJoltage = findMaxJoltagePart2(line)
				} else {
					maxJoltage = findMaxJoltage(line)
				}
				totalJoltage += maxJoltage
			}

			if totalJoltage != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, totalJoltage)
			}
		})
	}
}
