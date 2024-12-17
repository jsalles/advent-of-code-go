package main

import (
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
	maze, _ := ParseInput(input)
	score, path := dijktras(maze)
	if part2 {
		return getUniqueCount(maze, path)
	}
	// solve part 1 here
	return score
}

type Position struct {
	x, y int
}

type step struct {
	co      Position
	lastDir Position
	score   int
	path    map[Position]int
}

type point struct {
	co      Position
	lastDir Position
}

var (
	Up    = Position{0, -1}
	Down  = Position{0, 1}
	Left  = Position{-1, 0}
	Right = Position{1, 0}
)

func dijktras(matrix [][]string) (int, map[Position]int) {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{Position{1, len(matrix) - 2}, Right, 0, make(map[Position]int)})

	visited := make(map[point]struct{})

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		currentNode := element.(step)

		if _, ok := visited[point{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if matrix[currentNode.co.y][currentNode.co.x] == "E" {
			return currentNode.score, currentNode.path
		}

		nextSteps := getNextSteps(currentNode, matrix, visited)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}

		visited[point{currentNode.co, currentNode.lastDir}] = struct{}{}
	}
	return -1, make(map[Position]int)
}

func ParseInput(input string) (board [][]string, currentPos Position) {
	for y, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		var newRow []string
		for x, cell := range line {
			newRow = append(newRow, string(cell))

			if cell == 'S' {
				currentPos = Position{x, y}
			}
		}
		board = append(board, newRow)
	}

	return board, currentPos
}

func isValidStep(current Position, input [][]string) bool {
	if current.x < 0 || current.y < 0 || current.x >= len(input[0]) || current.y >= len(input) {
		return false
	}
	return true
}

func copyMap(path map[Position]int) map[Position]int {
	new := make(map[Position]int, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}

func getAllowedDirections(direction Position) []Position {
	switch direction {
	case Up:
		return []Position{Up, Left, Right}
	case Down:
		return []Position{Down, Left, Right}
	case Left:
		return []Position{Up, Left, Down}
	case Right:
		return []Position{Up, Down, Right}
	}
	return []Position{}
}

func getNextSteps(current step, grid [][]string, visited map[point]struct{}) []step {
	possibleNext := []step{}
	for _, dir := range getAllowedDirections(current.lastDir) {
		newPosition := Position{current.co.x + dir.x, current.co.y + dir.y}

		if !isValidStep(newPosition, grid) {
			continue
		}

		if grid[newPosition.y][newPosition.x] == "#" {
			continue
		}

		if _, ok := visited[point{newPosition, dir}]; ok {
			continue
		}

		score := current.score + 1
		if dir != current.lastDir {
			score += 1000
		}

		possibleNext = append(possibleNext, step{
			co:      newPosition,
			lastDir: dir,
			score:   score,
			path:    copyMap(current.path),
		})
	}
	return possibleNext
}

func getUniqueCount(matrix [][]string, path map[Position]int) int {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(step).score - b.(step).score
	})

	priorityQueue.Enqueue(step{Position{1, len(matrix) - 2}, Right, 0, make(map[Position]int)})

	visited := make(map[point]struct{})
	newSafeCoordinates := make(map[Position]struct{})

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()

		currentNode := element.(step)

		if score, ok := path[currentNode.co]; ok && score == currentNode.score {
			for point := range currentNode.path {
				if _, ok := path[point]; !ok {
					newSafeCoordinates[point] = struct{}{}
				}
			}
		}

		if _, ok := visited[point{currentNode.co, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.co] = currentNode.score

		if matrix[currentNode.co.y][currentNode.co.x] == "E" {
			continue
		}

		nextSteps := getNextSteps(currentNode, matrix, visited)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}

		visited[point{currentNode.co, currentNode.lastDir}] = struct{}{}
	}
	return len(path) + len(newSafeCoordinates)
}
