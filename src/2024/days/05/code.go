package main

import (
	"fmt"
	"slices"
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
	ordering, sequences := parse_input(input)
	sum := 0

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, sequence := range sequences {
			if !is_valid_sequence(sequence, ordering) {
				fixed_sequence := sort_list(sequence, ordering)
				if !is_valid_sequence(fixed_sequence, ordering) {
					fmt.Println("THIS IS WRONG")
					return -1
				}
				center, _ := strconv.Atoi(fixed_sequence[len(fixed_sequence)/2])
				sum += center
			}
		}
		return sum
	}
	// solve part 1 here
	for _, sequence := range sequences {
		if is_valid_sequence(sequence, ordering) {
			center, _ := strconv.Atoi(sequence[len(sequence)/2])
			sum += center
		}
	}
	return sum
}

func find_index(slice []string, element string) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

func sort_list(sequence []string, ordering_rules map[string]Ordering) []string {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(sequence)-1; i++ {
			if slices.Contains(ordering_rules[sequence[i]].smaller, sequence[i+1]) {
				sequence[i], sequence[i+1] = sequence[i+1], sequence[i]
				swapped = true
			}
		}
	}

	return sequence
}

func is_valid_sequence(sequence []string, ordering_rules map[string]Ordering) bool {
	for i := 1; i < len(sequence); i++ {
		if slices.Contains(ordering_rules[sequence[i]].larger, sequence[i-1]) {
			return false
		}
	}
	return true
}

type Ordering struct {
	smaller []string
	larger  []string
}

func parse_input(input string) (map[string]Ordering, [][]string) {
	ordering := make(map[string]Ordering)
	sequences := make([][]string, 0)

	blocks := strings.Split(input, "\n\n")
	for _, ordering_rule := range strings.Split(blocks[0], "\n") {
		numbers := strings.Split(ordering_rule, "|")
		before, after := numbers[0], numbers[1]
		_, exists := ordering[before]
		if !exists {
			ordering[before] = Ordering{smaller: make([]string, 0), larger: make([]string, 0)}
		}
		temp := ordering[before]
		temp.larger = append(temp.larger, after)
		ordering[before] = temp

		_, exists = ordering[after]
		if !exists {
			ordering[after] = Ordering{smaller: make([]string, 0), larger: make([]string, 0)}
		}
		temp = ordering[after]
		temp.smaller = append(temp.smaller, before)
		ordering[after] = temp
	}

	for _, sequence := range strings.Split(blocks[1], "\n") {
		sequences = append(sequences, strings.Split(sequence, ","))
	}

	return ordering, sequences
}
