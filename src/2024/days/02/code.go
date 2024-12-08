package main

import (
	arrutils "aoc-in-go/src/2024/utils/array_utils"
	"math"
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
	number_rows := parseInput(input)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		count := 0
		for _, numbers := range number_rows {
			if isValidSequence(numbers) {
				count++
			} else {
				for i := 0; i < len(numbers); i++ {
					new_numbers := make([]int, 0)
					new_numbers = append(new_numbers, numbers[:i]...)
					new_numbers = append(new_numbers, numbers[i+1:]...)

					if isValidSequence(new_numbers) {
						count++
						break
					}
				}
			}
		}
		return count
	}

	// solve part 1 here
	count := 0
	for _, numbers := range number_rows {
		if isValidSequence(numbers) {
			count++
		}
	}
	return count
}

func isValidSequence(numbers []int) bool {
	ascending := numbers[1] > numbers[0]
	for i := 1; i < len(numbers); i++ {
		if ascending && numbers[i] <= numbers[i-1] || !ascending && numbers[i] >= numbers[i-1] || math.Abs(float64(numbers[i]-numbers[i-1])) > 3 {
			return false
		}
	}
	return true
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	parsedInput := make([][]int, len(lines))
	for i, line := range lines {
		numbers := arrutils.Map(strings.Fields(line), func(s string) int {
			num, _ := strconv.Atoi(s)
			return num
		})
		parsedInput[i] = numbers
	}
	return parsedInput
}
