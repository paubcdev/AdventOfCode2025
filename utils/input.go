package utils

import (
	"bufio"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadNonEmptyLines(filename string) ([]string, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var nonEmpty []string
	for _, line := range lines {
		if line != "" {
			nonEmpty = append(nonEmpty, line)
		}
	}

	return nonEmpty, nil
}
