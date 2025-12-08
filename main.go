package main

import (
	"fmt"
	"os"
	"strconv"

	"AoC2025/days/day01"
	"AoC2025/days/day02"
	"AoC2025/days/day03"
	"AoC2025/days/day04"
	"AoC2025/days/day05"
	"AoC2025/days/day06"
	"AoC2025/days/day07"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day> [part]")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 12 {
		fmt.Printf("Invalid input: %s (day must be between 1 - 12)\n", os.Args[1])
		os.Exit(1)
	}

	part := 0
	if len(os.Args) >= 3 {
		part, err = strconv.Atoi(os.Args[2])
		if err != nil || (part != 1 && part != 2) {
			fmt.Printf("Invalid part: %s (must be 1 or 2)\n", os.Args[2])
			os.Exit(1)
		}
	}

	fmt.Printf("=== Advent of Code 2025 - Day %d ===\n\n", day)

	solver := getSolver(day)
	if solver == nil {
		fmt.Printf("Day %d not available\n", day)
		os.Exit(1)
	}

	solver.Run(part)
}

type Solver interface {
	Run(part int)
}

func getSolver(day int) Solver {
	switch day {
	case 1:
		return &day01.Solution{}
	case 2:
		return &day02.Solution{}
	case 3:
		return &day03.Solution{}
	case 4:
		return &day04.Solution{}
	case 5:
		return &day05.Solution{}
	case 6:
		return &day06.Solution{}
	case 7:
		return &day07.Solution{}
	default:
		return nil
	}
}
