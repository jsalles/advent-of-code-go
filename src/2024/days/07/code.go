package main

import (
	"fmt"
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
	problems := ParseInput(input)

	var backtrackFunc func(int, []int, int) bool
	backtrackFunc = func(expected int, numbers []int, current int) bool {
		if current > expected {
			return false
		}
		if len(numbers) == 0 {
			return current == expected
		}

		if backtrackFunc(expected, numbers[1:], current+numbers[0]) {
			return true
		}
		if part2 {
			joined, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, numbers[0]))
			if backtrackFunc(expected, numbers[1:], joined) {
				return true
			}
		}
		return backtrackFunc(expected, numbers[1:], current*numbers[0])
	}

	// solve part 1 here
	sum := 0
	for _, problem := range problems {
		if backtrackFunc(problem.expected, problem.numbers, 0) {
			sum += problem.expected
		}
	}
	// problem := problems[4]
	// backtrackFunc(problem.expected, problem.numbers, 0)

	return sum
}

type Problem struct {
	numbers  []int
	expected int
}

func ParseInput(input string) []Problem {
	var problems []Problem
	for _, line := range strings.Split(input, "\n") {
		problem := Problem{}
		parts := strings.Split(line, ":")
		expected, _ := strconv.Atoi(parts[0])
		problem.expected = expected
		problem.numbers = make([]int, 0)

		for _, number := range strings.Split(strings.Trim(parts[1], " "), " ") {
			parsedNumber, _ := strconv.Atoi(number)
			problem.numbers = append(problem.numbers, parsedNumber)
		}

		problems = append(problems, problem)
	}

	return problems
}
