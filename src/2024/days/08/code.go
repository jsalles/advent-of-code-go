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
	grid, height, width := ParseInput(input)
	antinodeMap := make(map[Point]bool, 0)
	for i := 0; i < len(grid)-1; i++ {
		for j := i + 1; j < len(grid); j++ {
			if grid[i].value == grid[j].value {
				deltaRow := grid[i].point.row - grid[j].point.row
				deltaCol := grid[i].point.col - grid[j].point.col

				pointsAbove := AddAntinodes(grid[i].point, Point{row: deltaRow, col: deltaCol}, Point{row: height, col: width}, part2)
				for _, pointAbove := range pointsAbove {
					antinodeMap[pointAbove] = true
				}

				pointsBelow := AddAntinodes(grid[i].point, Point{row: -deltaRow, col: -deltaCol}, Point{row: height, col: width}, part2)
				for _, pointBelow := range pointsBelow {
					antinodeMap[pointBelow] = true
				}

				antinodeMap[grid[i].point] = true
				antinodeMap[grid[j].point] = true
			}
		}
	}

	return len(antinodeMap)
}

type Point struct {
	row, col int
}

type Cell struct {
	value string
	point Point
}

func AddAntinodes(point Point, delta Point, dims Point, repeat bool) []Point {
	antinodes := make([]Point, 0)

	candidate := Point{row: point.row + delta.row, col: point.col + delta.col}
	for IsInGrid(candidate, dims.row, dims.col) {
		antinodes = append(antinodes, candidate)
		if !repeat {
			break
		}

		candidate = Point{row: candidate.row + delta.row, col: candidate.col + delta.col}
	}

	return antinodes
}

func IsInGrid(point Point, height int, width int) bool {
	return point.row >= 0 && point.row < height && point.col >= 0 && point.col < width
}

func ParseInput(input string) ([]Cell, int, int) {
	grid := make([]Cell, 0)
	for row, line := range strings.Split(input, "\n") {
		for col, cell := range strings.Split(line, "") {
			if cell == "." {
				continue
			}
			point := Point{row, col}
			grid = append(grid, Cell{point: point, value: cell})
		}
	}

	height := strings.Count(strings.Trim(input, "\n"), "\n") + 1
	width := strings.Index(input, "\n")
	return grid, height, width
}
