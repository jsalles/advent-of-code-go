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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	keys, locks := parseInput(input)
	var locksPinHeights, keysPinHeights [][5]int
	for _, key := range keys {
		keysPinHeights = append(keysPinHeights, getKeyPinHeights(key))
	}
	for _, lock := range locks {
		locksPinHeights = append(locksPinHeights, getLockPinHeights(lock))
	}

	sum := 0
	for _, lock := range locksPinHeights {
		for _, key := range keysPinHeights {
			fits := true
			for i := range 5 {
				if lock[i]+key[i] > 5 {
					fits = false
					break
				}
			}
			if fits {
				sum++
			}
		}
	}

	return sum
}

func getLockPinHeights(pattern Pattern) [5]int {
	var pinHeights [5]int

	for x := 0; x < 5; x++ {
		maxHeight := 6
		for pattern[maxHeight][x] == "." {
			maxHeight--
		}
		pinHeights[x] = maxHeight
	}

	return pinHeights
}

func getKeyPinHeights(pattern Pattern) [5]int {
	var pinHeights [5]int

	for x := 0; x < 5; x++ {
		minHeight := 0
		for pattern[minHeight+1][x] == "." {
			minHeight++
		}
		pinHeights[x] = 5 - minHeight
	}

	return pinHeights
}

type Pattern [7][5]string

func parseInput(input string) ([]Pattern, []Pattern) {
	var keys, locks []Pattern
	for _, block := range strings.Split(input, "\n\n") {
		var pattern Pattern
		for y, line := range strings.Split(block, "\n") {
			for x, cell := range line {
				pattern[y][x] = string(cell)
			}
		}
		if pattern[0][0] == "#" {
			locks = append(locks, pattern)
		} else {
			keys = append(keys, pattern)
		}
	}

	return keys, locks
}
