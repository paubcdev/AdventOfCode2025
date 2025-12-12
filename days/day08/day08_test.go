package day08

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay08(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday08")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	points := parsePoints(lines)

	tests := []struct {
		part        int
		connections int
		expected    int
	}{
		{1, 10, 40},
		{2, 0, 25272},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.part), func(t *testing.T) {
			var result int
			if tt.part == 1 {
				result = connectJunctionBoxes(points, tt.connections)
			} else {
				result = findLastConnection(points)
			}

			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
