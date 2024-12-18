package main

import (
	"fmt"
	"strconv"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
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
	grid, extraBytes := parseInput(input)
	minSteps, pathMap := dijkstras(grid)
	if part2 {
		firstBreakingByte := findFirstBreakingByte(grid, pathMap, extraBytes)
		return firstBreakingByte
	}
	// solve part 1 here

	return minSteps
}

func findFirstBreakingByte(grid [][]bool, pathMap map[Position]bool, extraBytes []Position) Position {
	for _, pos := range extraBytes {
		grid[pos.y][pos.x] = true
		if !pathMap[pos] {
			continue
		}

		count, newPath := dijkstras(grid)
		if count == -1 {
			return pos
		}
		pathMap = newPath
	}

	return Position{}
}

type Position struct {
	x, y int
}

type Step struct {
	pos   Position
	score int
}

func dijkstras(grid [][]bool) (int, map[Position]bool) {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})

	priorityQueue.Enqueue(Step{Position{0, 0}, 0})
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := []Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	predecessorMap := make(map[Position]Position)

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()
		current := element.(Step)
		x, y := current.pos.x, current.pos.y

		if visited[y][x] {
			continue
		}

		visited[y][x] = true

		if y == rows-1 && x == cols-1 {
			path := make(map[Position]bool)
			pos := current.pos
			for pos.x != 0 && pos.y != 0 {
				path[pos] = true
				pos = predecessorMap[pos]
			}
			return current.score, path
		}

		for _, dir := range directions {
			newPos := Position{current.pos.x + dir.x, current.pos.y + dir.y}
			if isInGrid(newPos, grid) && !visited[newPos.y][newPos.x] && !grid[newPos.y][newPos.x] {
				predecessorMap[newPos] = current.pos
				priorityQueue.Enqueue(Step{newPos, current.score + 1})
			}
		}
	}

	return -1, map[Position]bool{}
}

func copyPath(path map[Position]bool) map[Position]bool {
	result := make(map[Position]bool, len(path))
	for k, v := range path {
		result[k] = v
	}
	return result
}

func isInGrid(pos Position, grid [][]bool) bool {
	return pos.y >= 0 && pos.y < len(grid) && pos.x >= 0 && pos.x < len(grid[0])
}

func parseInput(input string) ([][]bool, []Position) {
	blocks := strings.Split(input, "\n\n")
	parameters := strings.Split(blocks[0], "\n")
	bytesToRead, _ := strconv.Atoi(parameters[0])
	dim, _ := strconv.Atoi(parameters[1])
	dim += 1

	grid := make([][]bool, dim)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, dim)
	}

	for _, pair := range strings.Split(strings.Trim(blocks[1], "\n"), "\n")[:bytesToRead] {
		coords := strings.Split(pair, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = true
	}

	extraBytes := make([]Position, 0)
	for _, pair := range strings.Split(strings.Trim(blocks[1], "\n"), "\n")[bytesToRead:] {
		coords := strings.Split(pair, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		extraBytes = append(extraBytes, Position{x, y})
	}

	return grid, extraBytes
}

func printGrid(grid [][]bool, pathMap map[Position]bool) {
	for y, row := range grid {
		for x, cell := range row {
			if pathMap[Position{x, y}] {
				fmt.Print("O")
			} else if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
