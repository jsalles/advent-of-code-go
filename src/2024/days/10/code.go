package main

import (
	"container/list"
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
	grid, zeros := ParseInput(input)
	sum := 0
	for _, zeroPos := range zeros {
		sum += bfs(grid, zeroPos, !part2)
	}
	return sum
}

var height, width int

type Position struct {
	row, col int
}

type Cell struct {
	pos   Position
	value int
}

func bfs(grid [][]int, start Position, countUnique bool) int {
	directions := [4]Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	seen := make(map[Position]bool, 0)

	queue := list.New()
	queue.PushBack(Cell{start, 0})

	sum := 0
	for queue.Len() > 0 {
		queueElement := queue.Back()
		queue.Remove(queueElement)
		currentCell, ok := queueElement.Value.(Cell)
		if !ok {
			panic("oops")
		}

		for _, direction := range directions {
			newRow := currentCell.pos.row + direction.row
			newCol := currentCell.pos.col + direction.col

			newPos := Position{newRow, newCol}
			if !IsInGrid(newPos) || !IsUphillSlope(currentCell.value, grid[newRow][newCol]) {
				continue
			}
			if grid[newRow][newCol] == 9 {
				if countUnique && seen[newPos] {
					continue
				}
				sum++
				seen[newPos] = true
				continue
			}

			queue.PushBack(Cell{newPos, grid[newRow][newCol]})
		}
	}

	return sum
}

func IsInGrid(position Position) bool {
	return position.row >= 0 && position.row < height && position.col >= 0 && position.col < width
}

func IsUphillSlope(start int, end int) bool {
	return end-start == 1
}

func ParseInput(input string) ([][]int, []Position) {
	height = strings.Count(strings.Trim(input, "\n"), "\n") + 1
	width = strings.Index(input, "\n")

	grid := make([][]int, height)
	zeros := make([]Position, 0)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}

	for row, line := range strings.Split(input, "\n") {
		for col, ch := range line {
			num := int(ch - '0')

			grid[row][col] = num
			if num == 0 {
				zeros = append(zeros, Position{row, col})
			}
		}
	}

	return grid, zeros
}
