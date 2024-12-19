package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	patterns, designs := parseInput(input)
	sum := 0
	mem := make(map[string]int)
	for _, design := range designs {
		possibilities := designPossibilities(design, patterns, mem)
		if part2 {
			sum += possibilities
		} else if possibilities > 0 {
			sum += 1
		}
	}
	return sum
}

func designPossibilities(design string, patterns []string, mem map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	if val, exists := mem[design]; exists {
		return val
	}

	count := 0
	for _, pattern := range patterns {
		patternLength := len(pattern)
		if patternLength <= len(design) && design[:patternLength] == pattern {
			count += designPossibilities(design[patternLength:], patterns, mem)
		}

	}
	mem[design] = count
	return count
}

func parseInput(input string) ([]string, []string) {
	blocks := strings.Split(input, "\n\n")
	patterns := make([]string, 0)
	for _, pattern := range strings.Split(blocks[0], ",") {
		patterns = append(patterns, strings.Trim(pattern, " "))
	}

	designs := make([]string, 0)
	for _, design := range strings.Split(blocks[1], "\n") {
		designs = append(designs, design)
	}

	return patterns, designs
}
