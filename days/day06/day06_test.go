package day06

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay06(t *testing.T) {
	tests := []struct {
		part     int
		expected int64
	}{
		{part: 1, expected: 4277556},
		{part: 2, expected: 3263827},
	}

	for _, tt := range tests {
		t.Run("Part"+strconv.Itoa(tt.part), func(t *testing.T) {
			lines, err := utils.ReadLines("../../test_inputs/testday06")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			var grandTotal int64
			if tt.part == 2 {
				grandTotal = solveWorksheetPart2(lines)
			} else {
				grandTotal = solveWorksheet(lines)
			}

			if grandTotal != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, grandTotal)
			}
		})
	}
}
