package day03

import (
	"fmt"
	"strconv"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day03")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	totalJoltage := int64(0)

	for _, line := range lines {
		var maxJoltage int64
		if part == 2 {
			maxJoltage = findMaxJoltagePart2(line)
		} else {
			maxJoltage = findMaxJoltage(line)
		}
		totalJoltage += maxJoltage
	}

	if part == 2 {
		fmt.Printf("Part 2: %d\n", totalJoltage)
	} else {
		fmt.Printf("Part 1: %d\n", totalJoltage)
	}
}

func findMaxJoltage(bank string) int64 {
	maxJoltage := int64(0)

	// Try all pairs of positions
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Form the number from batteries at positions i and j
			joltageStr := string(bank[i]) + string(bank[j])
			joltage, err := strconv.ParseInt(joltageStr, 10, 64)
			if err != nil {
				continue
			}

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}

	return maxJoltage
}

func findMaxJoltagePart2(bank string) int64 {
	n := len(bank)
	if n < 12 {
		return 0
	}

	// Greedy approach: to maximize a 12-digit number,
	// at each position, pick the largest available digit
	// We need to skip (n - 12) digits total

	toSkip := n - 12
	result := ""
	i := 0

	for len(result) < 12 && i < n {
		// Look ahead to see the best digit we can pick
		// We can skip at most (toSkip) more digits
		maxDigit := byte('0')
		maxPos := i

		// Check the next (toSkip + 1) positions
		for j := i; j <= i+toSkip && j < n; j++ {
			if bank[j] > maxDigit {
				maxDigit = bank[j]
				maxPos = j
			}
		}

		// Pick the maximum digit found
		result += string(maxDigit)
		// Update how many we can skip and where we are
		toSkip -= (maxPos - i)
		i = maxPos + 1
	}

	joltage, _ := strconv.ParseInt(result, 10, 64)
	return joltage
}
