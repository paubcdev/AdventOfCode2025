package day10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

type Machine struct {
	target   []bool
	buttons  [][]int
	joltages []int
}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day10")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	machines := parseMachines(lines)
	totalPresses := 0

	if part == 1 {
		for _, machine := range machines {
			presses := solveGaussian(machine)
			totalPresses += presses
		}
		fmt.Printf("Part 1: Total button presses = %d\n", totalPresses)
	} else {
		for _, machine := range machines {
			presses := solveJoltage(machine)
			totalPresses += presses
		}
		fmt.Printf("Part 2: Total button presses = %d\n", totalPresses)
	}
}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, 0, len(lines))

	for _, line := range lines {
		machine := parseMachine(line)
		machines = append(machines, machine)
	}

	return machines
}

func parseMachine(line string) Machine {
	bracketRe := regexp.MustCompile(`\[([.#]+)\]`)
	parenRe := regexp.MustCompile(`\(([0-9,]+)\)`)
	braceRe := regexp.MustCompile(`\{([0-9,]+)\}`)

	bracketMatch := bracketRe.FindStringSubmatch(line)
	parenMatches := parenRe.FindAllStringSubmatch(line, -1)
	braceMatch := braceRe.FindStringSubmatch(line)

	target := make([]bool, 0)
	if len(bracketMatch) > 1 {
		for _, ch := range bracketMatch[1] {
			target = append(target, ch == '#')
		}
	}

	buttons := make([][]int, 0)
	for _, match := range parenMatches {
		if len(match) > 1 {
			indices := make([]int, 0)
			parts := strings.Split(match[1], ",")
			for _, part := range parts {
				idx, _ := strconv.Atoi(strings.TrimSpace(part))
				indices = append(indices, idx)
			}
			buttons = append(buttons, indices)
		}
	}

	joltages := make([]int, 0)
	if len(braceMatch) > 1 {
		parts := strings.Split(braceMatch[1], ",")
		for _, part := range parts {
			val, _ := strconv.Atoi(strings.TrimSpace(part))
			joltages = append(joltages, val)
		}
	}

	return Machine{target: target, buttons: buttons, joltages: joltages}
}

func solveGaussian(m Machine) int {
	n := len(m.target)
	numButtons := len(m.buttons)

	if numButtons == 0 {
		for _, t := range m.target {
			if t {
				return 1 << 30
			}
		}
		return 0
	}

	minPresses := 1 << 30

	for mask := 0; mask < (1 << numButtons); mask++ {
		state := make([]bool, n)
		presses := 0

		for b := 0; b < numButtons; b++ {
			if (mask & (1 << b)) != 0 {
				presses++
				for _, idx := range m.buttons[b] {
					if idx < n {
						state[idx] = !state[idx]
					}
				}
			}
		}

		match := true
		for i := 0; i < n; i++ {
			if state[i] != m.target[i] {
				match = false
				break
			}
		}

		if match && presses < minPresses {
			minPresses = presses
		}
	}

	if minPresses == 1<<30 {
		return 0
	}

	return minPresses
}

func solveJoltage(m Machine) int {
	numButtons := len(m.buttons)
	numCounters := len(m.joltages)

	if numButtons == 0 {
		for _, j := range m.joltages {
			if j != 0 {
				return 1 << 30
			}
		}
		return 0
	}

	matrix := make([][]int, numCounters)
	for i := 0; i < numCounters; i++ {
		matrix[i] = make([]int, numButtons+1)
		matrix[i][numButtons] = m.joltages[i]
	}

	for j, button := range m.buttons {
		for _, idx := range button {
			if idx < numCounters {
				matrix[idx][j] = 1
			}
		}
	}

	pivotRow := 0
	pivotCols := make([]int, 0)

	for col := 0; col < numButtons && pivotRow < numCounters; col++ {
		found := -1
		for r := pivotRow; r < numCounters; r++ {
			if matrix[r][col] != 0 {
				found = r
				break
			}
		}

		if found == -1 {
			continue
		}

		if found != pivotRow {
			matrix[pivotRow], matrix[found] = matrix[found], matrix[pivotRow]
		}

		pivotCols = append(pivotCols, col)

		for r := 0; r < numCounters; r++ {
			if r == pivotRow || matrix[r][col] == 0 {
				continue
			}

			factor := matrix[r][col]
			pivot := matrix[pivotRow][col]

			for c := 0; c <= numButtons; c++ {
				matrix[r][c] = matrix[r][c]*pivot - matrix[pivotRow][c]*factor
			}
		}

		pivotRow++
	}

	for r := pivotRow; r < numCounters; r++ {
		if matrix[r][numButtons] != 0 {
			return 0
		}
	}

	isFree := make([]bool, numButtons)
	for i := 0; i < numButtons; i++ {
		isFree[i] = true
	}
	for _, col := range pivotCols {
		isFree[col] = false
	}

	maxJoltage := 0
	for _, j := range m.joltages {
		if j > maxJoltage {
			maxJoltage = j
		}
	}

	minPresses := 1 << 30

	freeVars := make([]int, 0)
	for i := 0; i < numButtons; i++ {
		if isFree[i] {
			freeVars = append(freeVars, i)
		}
	}

	numFree := len(freeVars)

	var search func(freeIdx int, freeVals []int)
	search = func(freeIdx int, freeVals []int) {
		if freeIdx == numFree {
			solution := make([]int, numButtons)
			for i, val := range freeVals {
				solution[freeVars[i]] = val
			}

			valid := true
			for i := len(pivotCols) - 1; i >= 0; i-- {
				col := pivotCols[i]
				pivot := matrix[i][col]

				if pivot == 0 {
					valid = false
					break
				}

				rhs := matrix[i][numButtons]
				for j := 0; j < numButtons; j++ {
					if j != col {
						rhs -= matrix[i][j] * solution[j]
					}
				}

				if rhs%pivot != 0 {
					valid = false
					break
				}

				solution[col] = rhs / pivot
				if solution[col] < 0 {
					valid = false
					break
				}
			}

			if valid {
				achieved := make([]int, numCounters)
				for btnIdx := 0; btnIdx < numButtons; btnIdx++ {
					for _, counterIdx := range m.buttons[btnIdx] {
						if counterIdx < numCounters {
							achieved[counterIdx] += solution[btnIdx]
						}
					}
				}

				verified := true
				for i := 0; i < numCounters; i++ {
					if achieved[i] != m.joltages[i] {
						verified = false
						break
					}
				}

				if verified {
					sum := 0
					for _, p := range solution {
						sum += p
					}
					if sum < minPresses {
						minPresses = sum
					}
				}
			}
			return
		}

		for val := 0; val <= maxJoltage; val++ {
			newFreeVals := make([]int, len(freeVals)+1)
			copy(newFreeVals, freeVals)
			newFreeVals[len(freeVals)] = val

			partialSum := 0
			for _, p := range newFreeVals {
				partialSum += p
			}
			if partialSum >= minPresses {
				break
			}

			search(freeIdx+1, newFreeVals)
		}
	}

	search(0, []int{})

	if minPresses == 1<<30 {
		return 0
	}

	return minPresses
}
