package main

import (
	"fmt"
	"strconv"

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
	pattern, patternArrangement := ParseInput(input)
	if part2 {
		for filledIndex := len(patternArrangement.filled) - 1; filledIndex >= 0; filledIndex-- {
			filled := patternArrangement.filled[filledIndex]
			lengthFilled := filled.end - filled.start + 1
			for emptyIndex := 0; emptyIndex < len(patternArrangement.empty); emptyIndex++ {
				empty := patternArrangement.empty[emptyIndex]
				if empty.start > filled.end {
					break
				}
				lengthEmpty := empty.end - empty.start + 1

				if lengthFilled <= lengthEmpty {
					for i := 0; i < lengthFilled; i++ {
						pattern[i+empty.start] = Cell{value: filled.value, isEmpty: false}
					}
					for i := filled.start; i <= filled.end; i++ {
						pattern[i] = Cell{isEmpty: true}
					}

					if lengthEmpty == lengthFilled {
						patternArrangement.empty = append(patternArrangement.empty[:emptyIndex], patternArrangement.empty[emptyIndex+1:]...)
					} else {
						patternArrangement.empty[emptyIndex].start += lengthFilled
					}
					break
				}
			}
		}
		return SumPattern(pattern)
	}
	// solve part 1 here
	left, right := GetInitialPointers(pattern)
	for left < right {
		pattern[left] = Cell{value: pattern[right].value, isEmpty: false}
		pattern[right].isEmpty = true

		for !pattern[left].isEmpty {
			left++
		}
		for pattern[right].isEmpty {
			right--
		}
	}

	return SumPattern(pattern)
}

type Cell struct {
	value   int
	isEmpty bool
}

type Arrangement struct {
	start, end, value int
}

type PatternArrangement struct {
	empty  []Arrangement
	filled []Arrangement
}

func SumPattern(pattern []Cell) int {
	sum := 0
	for i, cell := range pattern {
		if !cell.isEmpty {
			sum += cell.value * i
		}
		i++
	}

	return sum
}

func GetInitialPointers(pattern []Cell) (int, int) {
	left, right := 0, len(pattern)-1
	for !pattern[left].isEmpty {
		left++
	}
	for pattern[right].isEmpty {
		right--
	}
	return left, right
}

func ParseInput(input string) ([]Cell, PatternArrangement) {
	var pattern []Cell
	patternArrangement := PatternArrangement{empty: make([]Arrangement, 0), filled: make([]Arrangement, 0)}
	isActive := true
	currentValue := 0

	// Process each number in the input
	for _, ch := range input {
		count, _ := strconv.Atoi(string(ch))

		// Add cells based on the current state
		cells := make([]Cell, count)
		for i := range cells {
			cells[i] = Cell{value: currentValue, isEmpty: !isActive}
		}
		pattern = append(pattern, cells...)

		// Update state
		arrangement := Arrangement{
			start: len(pattern) - count,
			end:   len(pattern) - 1,
			value: currentValue,
		}
		if isActive {
			patternArrangement.filled = append(patternArrangement.filled, arrangement)
			currentValue++
		} else {
			patternArrangement.empty = append(patternArrangement.empty, arrangement)
		}
		isActive = !isActive
	}

	return pattern, patternArrangement
}

func PrintPattern(pattern []Cell) {
	for _, cell := range pattern {
		if cell.isEmpty {
			fmt.Print(".")
		} else {
			fmt.Print(cell.value)
		}
	}
	fmt.Println()
	fmt.Println()
}
