package main

import (
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func find_digits(line string) int {
	left := -1
	var right int

	for _, c := range line {
		if !unicode.IsDigit(c) {
			continue
		}

		if left == -1 {
			left = int(c - '0')
		}
		right = int(c - '0')
	}

	return left*10 + right
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	substitutions := [][2]string{
		{"one", "on1e"},
		{"two", "tw2o"},
		{"three", "thre3e"},
		{"four", "fou4r"},
		{"five", "fiv5e"},
		{"six", "si6x"},
		{"seven", "seve7n"},
		{"eight", "eigh8t"},
		{"nine", "nin9e"},
	}
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		if part2 {
			for _, sub := range substitutions {
				line = strings.ReplaceAll(line, sub[0], sub[1])
			}
		}
		sum += find_digits(line)
	}

	return sum
}
