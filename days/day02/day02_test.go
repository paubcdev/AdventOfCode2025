package day02

import (
	"testing"

	"AoC2025/utils"
)

func TestDay02(t *testing.T) {
	tests := []struct {
		part     int
		expected int64
	}{
		{part: 1, expected: 1227775554},
		{part: 2, expected: 4174379265},
	}

	for _, tt := range tests {
		t.Run("Part"+string(rune('0'+tt.part)), func(t *testing.T) {
			lines, err := utils.ReadNonEmptyLines("../../test_inputs/testday02")
			if err != nil {
				t.Fatalf("Error reading test input: %v", err)
			}

			if len(lines) == 0 {
				t.Fatal("No input data")
			}

			ranges := parseRanges(lines[0])
			sum := int64(0)

			for _, r := range ranges {
				for id := r.start; id <= r.end; id++ {
					var invalid bool
					if tt.part == 2 {
						invalid = isInvalidIDPart2(id)
					} else {
						invalid = isInvalidID(id)
					}
					if invalid {
						sum += id
					}
				}
			}

			if sum != tt.expected {
				t.Errorf("Part %d: expected %d, got %d", tt.part, tt.expected, sum)
			}
		})
	}
}
