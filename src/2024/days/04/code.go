package main

import (
	"slices"
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

func run(isPart2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if isPart2 {
		return part2(input)
	}

	return part1(input)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	// directions going clockwise
	directions := [4][2]int{
		{-1, -1}, // top left
		{-1, 1},  // top right
		{1, 1},   // bottom right
		{1, -1},  // bottom left
	}
	valid_combinations := [4]string{"MSSM", "MMSS", "SMMS", "SSMM"}
	count := 0
	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			if lines[row][col] == 'A' {
				letters := ""

				for _, direction := range directions {
					newRow := row + direction[0]
					newCol := col + direction[1]
					letter := string(lines[newRow][newCol])

					letters += letter
				}

				if slices.Contains(valid_combinations[:], letters) {
					count++
				}
			}
		}
	}
	return count
}

func part1(input string) int {
	directions := [8][3][2]int{
		{{0, 1}, {0, 2}, {0, 3}},       // right
		{{0, -1}, {0, -2}, {0, -3}},    // left
		{{1, 0}, {2, 0}, {3, 0}},       // down
		{{-1, 0}, {-2, 0}, {-3, 0}},    // up
		{{-1, 1}, {-2, 2}, {-3, 3}},    // top right
		{{-1, -1}, {-2, -2}, {-3, -3}}, // top left
		{{1, 1}, {2, 2}, {3, 3}},       // bottom right
		{{1, -1}, {2, -2}, {3, -3}},    // bottom left
	}

	letters := [3]byte{'M', 'A', 'S'}
	// solve part 1 here
	lines := strings.Split(input, "\n")
	count := 0
	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			if line[col] == 'X' {
				for _, direction := range directions {
					found := true
					for i, cell := range direction {
						newRow := row + cell[0]
						newCol := col + cell[1]
						if newRow < 0 || newRow >= len(lines) || newCol < 0 || newCol >= len(line) {
							found = false
							break
						}
						if lines[newRow][newCol] != letters[i] {
							found = false
							break
						}

					}
					if found {
						count++
					}
				}
			}
		}
	}
	return count
}
