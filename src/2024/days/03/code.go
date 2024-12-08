package main

import (
	"regexp"
	"strconv"
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		parts := strings.Split(input, "do()")
		sum := 0
		for _, part := range parts {
			clean_part, _, _ := strings.Cut(part, "don't()")
			sum += sumMulsInLine(clean_part)
		}
		return sum
	}
	// solve part 1 here
	return sumMulsInLine(input)
}

func sumMulsInLine(line string) int {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindAllStringSubmatch(line, -1)

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	return sum
}
