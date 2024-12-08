package main

import (
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Position struct {
	row int
	col int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	board, startPosition := ParseInput(input)
	board[startPosition.row][startPosition.col] = "."
	directions := [4]Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	path := findPathToExit(board, startPosition, directions)

	if part2 {
		loopCount := findLoopCount(board, path, directions)
		return loopCount
	}

	return len(path)
}

func findLoopCount(board [][]string, path map[Position][]int, directions [4]Position) int {
	width := len(board[0])
	height := len(board)
	obstacleMap := make(map[Position]struct{})
	count := 0
	for step, stepDirections := range path {
		for _, direction := range stepDirections {
			// check if I'll create a loop, only if there isn't an obstacle directly in the next cell in my path
			nextPosition := Position{step.row + directions[direction].row, step.col + directions[direction].col}
			if !insideBoard(nextPosition, height, width) || board[nextPosition.row][nextPosition.col] == "#" {
				continue
			}

			board[nextPosition.row][nextPosition.col] = "#"
			foundLoop := isLoop(board, step, (direction+1)%len(directions), directions)
			board[nextPosition.row][nextPosition.col] = "."
			if foundLoop {
				if _, exists := obstacleMap[nextPosition]; !exists {
					count++
					obstacleMap[nextPosition] = struct{}{}
				}
				break
			}
		}
	}

	return count
}

func isLoop(board [][]string, guardPosition Position, directionIndex int, directions [4]Position) bool {
	width := len(board[0])
	height := len(board)
	visited := make(map[Position][]int)
	for insideBoard(guardPosition, height, width) {
		if board[guardPosition.row][guardPosition.col] == "#" {
			currentDirection := directions[directionIndex]
			// come back to previous position and turn
			newRow := guardPosition.row - currentDirection.row
			newCol := guardPosition.col - currentDirection.col
			guardPosition = Position{newRow, newCol}
			directionIndex = (directionIndex + 1) % len(directions)
		} else {
			seenDirections, exists := visited[guardPosition]
			if !exists {
				visited[guardPosition] = make([]int, 0)
			} else if slices.Contains(seenDirections, directionIndex) {
				return true
			}

			visited[guardPosition] = append(visited[guardPosition], directionIndex)

			// move to new position
			currentDirection := directions[directionIndex]
			newRow := guardPosition.row + currentDirection.row
			newCol := guardPosition.col + currentDirection.col
			guardPosition = Position{newRow, newCol}
		}
	}

	return false
}

func findPathToExit(board [][]string, guardPosition Position, directions [4]Position) map[Position][]int {
	currentDirectionIndex := 0
	width := len(board[0])
	height := len(board)
	visited := make(map[Position][]int)
	visited[guardPosition] = make([]int, 0)
	// visited[guardPosition] = append(visited[guardPosition], currentDirectionIndex)
	board[guardPosition.row][guardPosition.col] = "X"
	for insideBoard(guardPosition, height, width) {
		if board[guardPosition.row][guardPosition.col] == "#" {
			currentDirection := directions[currentDirectionIndex]
			// come back to previous position and turn
			newRow := guardPosition.row - currentDirection.row
			newCol := guardPosition.col - currentDirection.col
			guardPosition = Position{newRow, newCol}
			currentDirectionIndex = (currentDirectionIndex + 1) % len(directions)
		} else {
			// board[guardPosition.row][guardPosition.col] = "X"
			_, exists := visited[guardPosition]
			if exists {
			} else {
				visited[guardPosition] = make([]int, 0)
			}
			visited[guardPosition] = append(visited[guardPosition], currentDirectionIndex)

			// move to new position
			currentDirection := directions[currentDirectionIndex]
			newRow := guardPosition.row + currentDirection.row
			newCol := guardPosition.col + currentDirection.col
			guardPosition = Position{newRow, newCol}
		}
	}

	return visited
}

func insideBoard(position Position, height int, width int) bool {
	return position.row >= 0 &&
		position.row < height &&
		position.col >= 0 &&
		position.col < width
}

func ParseInput(input string) ([][]string, Position) {
	width := strings.Index(input, "\n")
	height := strings.Count(input, "\n")
	board := make([][]string, height+1)
	for i := 0; i < len(board); i++ {
		board[i] = make([]string, width)
	}
	var guardPosition Position

	for row, line := range strings.Split(input, "\n") {
		for col, cell := range strings.Split(line, "") {
			board[row][col] = cell

			if cell == "^" {
				guardPosition = Position{row, col}
			}
		}
	}

	return board, guardPosition
}
