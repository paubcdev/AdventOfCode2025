package day11

import (
"testing"

"AoC2025/utils"
)

func TestDay11(t *testing.T) {
	lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday11")
	if err != nil {
		t.Fatalf("Error reading test input: %v", err)
	}

	graph := parseGraph(lines)

	tests := []struct {
		part     int
		expected int
	}{
		{1, 5},
	}

	for _, tt := range tests {
		t.Run(string(rune('0'+tt.part)), func(t *testing.T) {
if tt.part == 1 {
count := countPaths(graph, "you", "out")
if count != tt.expected {
t.Errorf("Expected %d, got %d", tt.expected, count)
}
}
})
	}
}
