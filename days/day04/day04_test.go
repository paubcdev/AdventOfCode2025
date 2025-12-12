package day04

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay04(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{part: 1, expected: 13},
		{part: 2, expected: 43},
	}

	for _, tt := range tests {
		t.Run("Part"+strconv.Itoa(tt.part), func(t *testing.T) {
			lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday04")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			var result int
			if tt.part == 2 {
				result = removeAllAccessibleRolls(lines)
			} else {
				result = countAccessibleRolls(lines)
			}

			if result != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, result)
			}
		})
	}
}
