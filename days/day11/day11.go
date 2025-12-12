package day11

import (
	"fmt"
	"strings"

	"AoC2025/utils"
)

type Solution struct{}

func (s *Solution) Run(part int) {
	lines, err := utils.ReadNonEmptyLines("inputs/day11")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	graph := parseGraph(lines)

	if part == 1 {
		count := countPaths(graph, "you", "out")
		fmt.Printf("Part 1: %d\n", count)
	} else {
		count := countPathsWithRequired(graph, "svr", "out", []string{"dac", "fft"})
		fmt.Printf("Part 2: %d\n", count)
	}
}

func parseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		node := parts[0]
		outputs := strings.Fields(parts[1])
		graph[node] = outputs
	}

	return graph
}

func countPaths(graph map[string][]string, start, end string) int {
	if start == end {
		return 1
	}

	neighbors, exists := graph[start]
	if !exists {
		return 0
	}

	total := 0
	for _, neighbor := range neighbors {
		total += countPaths(graph, neighbor, end)
	}

	return total
}

func countPathsWithRequired(graph map[string][]string, start, end string, required []string) int {
	memo := make(map[string]int)
	return dfsWithMemo(graph, start, end, required, make(map[string]bool), make(map[string]bool), memo)
}

func dfsWithMemo(graph map[string][]string, current, end string, required []string, visited, seen map[string]bool, memo map[string]int) int {
	if current == end {
		for _, req := range required {
			if !seen[req] {
				return 0
			}
		}
		return 1
	}

	key := current + "|"
	for _, req := range required {
		if seen[req] {
			key += req
		}
	}
	if val, ok := memo[key]; ok {
		return val
	}

	neighbors, exists := graph[current]
	if !exists {
		return 0
	}

	visited[current] = true
	newSeen := make(map[string]bool)
	for k, v := range seen {
		newSeen[k] = v
	}
	for _, req := range required {
		if current == req {
			newSeen[req] = true
		}
	}

	total := 0
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			total += dfsWithMemo(graph, neighbor, end, required, visited, newSeen, memo)
		}
	}

	visited[current] = false
	memo[key] = total
	return total
}
