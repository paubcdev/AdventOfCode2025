package day10

import (
	"strconv"
	"testing"

	"AoC2025/utils"
)

func TestDay10(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday10")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	machines := parseMachines(lines)

	tests := []struct {
		part     int
		expected int
	}{
		{1, 7},
		{2, 33},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.part), func(t *testing.T) {
			totalPresses := 0
			for _, machine := range machines {
				var presses int
				if tt.part == 1 {
					presses = solveGaussian(machine)
				} else {
					presses = solveJoltage(machine)
				}
				totalPresses += presses
			}

			if totalPresses != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, totalPresses)
			}
		})
	}
}
