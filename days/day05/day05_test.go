package day05

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay05(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{part: 1, expected: 3},
		{part: 2, expected: 14},
	}

	for _, tt := range tests {
		t.Run("Part"+strconv.Itoa(tt.part), func(t *testing.T) {
			lines, err := utils.ReadLines("../../test_inputs/testday05")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			var result int
			if tt.part == 2 {
				result = countAllFreshIDs(lines)
			} else {
				result = countFreshIngredients(lines)
			}

			if result != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, result)
			}
		})
	}
}
