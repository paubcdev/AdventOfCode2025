package day06

import (
	"fmt"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadLines("inputs/day06")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var grandTotal int64
	if part == 2 {
		grandTotal = solveWorksheetPart2(lines)
		fmt.Printf("Part 2: %d\n", grandTotal)
	} else {
		grandTotal = solveWorksheet(lines)
		fmt.Printf("Part 1: %d\n", grandTotal)
	}
}

func solveWorksheet(lines []string) int64 {
	if len(lines) == 0 {
		return 0
	}

	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	for i := range lines {
		if len(lines[i]) < maxWidth {
			lines[i] = lines[i] + strings.Repeat(" ", maxWidth-len(lines[i]))
		}
	}

	grandTotal := int64(0)

	col := 0
	for col < maxWidth {
		if isEmptyColumn(lines, col) {
			col++
			continue
		}

		problem := extractProblem(lines, col)
		if problem != nil {
			result := solveProblem(problem)
			grandTotal += result
		}

		col++
		for col < maxWidth && !isEmptyColumn(lines, col) {
			col++
		}
	}

	return grandTotal
}

func isEmptyColumn(lines []string, col int) bool {
	for _, line := range lines {
		if col < len(line) && line[col] != ' ' {
			return false
		}
	}
	return true
}

type problem struct {
	numbers   []int64
	operation byte // '+' or '*'
}

func extractProblem(lines []string, startCol int) *problem {
	endCol := startCol
	for endCol < len(lines[0]) && !isEmptyColumn(lines, endCol) {
		endCol++
	}

	var numbers []int64
	var op byte

	for _, line := range lines {
		if startCol >= len(line) {
			continue
		}

		segment := ""
		if endCol <= len(line) {
			segment = strings.TrimSpace(line[startCol:endCol])
		} else {
			segment = strings.TrimSpace(line[startCol:])
		}

		if segment == "" {
			continue
		}

		if segment == "+" || segment == "*" {
			op = segment[0]
		} else {
			num, err := strconv.ParseInt(segment, 10, 64)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 {
		return nil
	}

	return &problem{
		numbers:   numbers,
		operation: op,
	}
}

func solveProblem(p *problem) int64 {
	if p == nil || len(p.numbers) == 0 {
		return 0
	}

	result := p.numbers[0]
	for i := 1; i < len(p.numbers); i++ {
		switch p.operation {
		case '+':
			result += p.numbers[i]
		case '*':
			result *= p.numbers[i]
		}
	}

	return result
}

func solveWorksheetPart2(lines []string) int64 {
	if len(lines) == 0 {
		return 0
	}

	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	for i := range lines {
		if len(lines[i]) < maxWidth {
			lines[i] = lines[i] + strings.Repeat(" ", maxWidth-len(lines[i]))
		}
	}

	grandTotal := int64(0)

	col := maxWidth - 1
	for col >= 0 {
		if isEmptyColumn(lines, col) {
			col--
			continue
		}

		problem := extractProblemPart2V2(lines, col)
		if problem != nil {
			result := solveProblemV2(problem)
			grandTotal += result
			col = problem.startCol - 1
		} else {
			col--
		}
	}

	return grandTotal
}

type problemV2 struct {
	numbers   []int64
	operation byte
	startCol  int
}

func extractProblemPart2V2(lines []string, endCol int) *problemV2 {
	startCol := endCol
	for startCol >= 0 && !isEmptyColumn(lines, startCol) {
		startCol--
	}
	startCol++

	var numbers []int64
	var op byte

	for col := endCol; col >= startCol; col-- {
		numStr := ""
		foundOp := byte(0)

		for row := 0; row < len(lines); row++ {
			if col >= len(lines[row]) {
				continue
			}

			ch := lines[row][col]
			if ch == ' ' {
				continue
			}

			if ch == '+' || ch == '*' {
				foundOp = ch
			} else if ch >= '0' && ch <= '9' {
				numStr += string(ch)
			}
		}

		if foundOp != 0 {
			op = foundOp
		}
		if numStr != "" {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 {
		return nil
	}

	return &problemV2{
		numbers:   numbers,
		operation: op,
		startCol:  startCol,
	}
}

func solveProblemV2(p *problemV2) int64 {
	if p == nil || len(p.numbers) == 0 {
		return 0
	}

	result := p.numbers[0]
	for i := 1; i < len(p.numbers); i++ {
		switch p.operation {
		case '+':
			result += p.numbers[i]
		case '*':
			result *= p.numbers[i]
		}
	}

	return result
}
