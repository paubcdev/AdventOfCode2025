package day02

import (
	"fmt"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day02")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	if len(lines) == 0 {
		fmt.Println("No input data")
		return
	}

	ranges := parseRanges(lines[0])
	sum := int64(0)

	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			if part == 2 {
				if isInvalidIDPart2(id) {
					sum += id
				}
			} else {
				if isInvalidID(id) {
					sum += id
				}
			}
		}
	}

	if part == 2 {
		fmt.Printf("Part 2: %d\n", sum)
	} else {
		fmt.Printf("Part 1: %d\n", sum)
	}
}

type idRange struct {
	start int64
	end   int64
}

func parseRanges(line string) []idRange {
	var ranges []idRange
	parts := strings.Split(line, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		rangeParts := strings.Split(part, "-")
		if len(rangeParts) != 2 {
			continue
		}

		start, err1 := strconv.ParseInt(rangeParts[0], 10, 64)
		end, err2 := strconv.ParseInt(rangeParts[1], 10, 64)

		if err1 == nil && err2 == nil {
			ranges = append(ranges, idRange{start: start, end: end})
		}
	}

	return ranges
}

func isInvalidID(id int64) bool {
	s := strconv.FormatInt(id, 10)
	n := len(s)

	// Must be even length to split in half
	if n%2 != 0 {
		return false
	}

	half := n / 2
	firstHalf := s[:half]
	secondHalf := s[half:]

	return firstHalf == secondHalf
}

func isInvalidIDPart2(id int64) bool {
	s := strconv.FormatInt(id, 10)
	n := len(s)

	// Try all possible divisors of the length
	for divisor := 2; divisor <= n; divisor++ {
		if n%divisor != 0 {
			continue
		}

		patternLen := n / divisor
		pattern := s[:patternLen]
		isAPattern := true

		for i := 0; i < divisor; i++ {
			start := i * patternLen
			end := start + patternLen
			if s[start:end] != pattern {
				isAPattern = false
				break
			}
		}

		if isAPattern {
			return true
		}
	}

	return false
}
