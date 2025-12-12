package day07

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay07(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday07")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	tests := []struct {
		part     int
		expected int
	}{
		{1, 21},
		{2, 40},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.part), func(t *testing.T) {
			var result int
			if tt.part == 1 {
				result = countBeamSplits(lines)
			} else {
				result = countTimelines(lines)
			}

			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
