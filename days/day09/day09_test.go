package day09

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay09(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday09")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	points := parsePoints(lines)

	tests := []struct {
		part     int
		expected int
	}{
		{1, 50},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.part), func(t *testing.T) {
			result := findLargestRectangle(points)

			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
