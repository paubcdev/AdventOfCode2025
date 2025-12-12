package day12

import (
	"testing"

	"AoC2025/utils"
)

func TestDay12(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday12")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	shapes, regions := parseInput(lines)

	tests := []struct {
		part     int
		expected int
	}{
		{1, 2},
	}

	for _, tt := range tests {
		t.Run(string(rune('0'+tt.part)), func(t *testing.T) {
			if tt.part == 1 {
				count := 0
				for _, region := range regions {
					if canFit(region, shapes) {
						count++
					}
				}
				if count != tt.expected {
					t.Errorf("Expected %d, got %d", tt.expected, count)
				}
			}
		})
	}
}
