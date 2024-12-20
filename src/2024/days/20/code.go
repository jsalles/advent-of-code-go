package main

import (
	"slices"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func ManhattanDistance(p1, p2 Position) int {
	dx := abs(p1.x - p2.x)
	dy := abs(p1.y - p2.y)
	return dx + dy
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	grid, start, _ := parseInput(input)
	_, path := dijkstras(grid, start)
	if part2 {
		count := 0
		for d1 := 0; d1 < len(path)-100; d1++ {
			for d2 := d1 + 101; d2 < len(path); d2++ {
				p1, p2 := path[d1], path[d2]
				dist := abs(p1.x-p2.x) + abs(p1.y-p2.y)
				if d2-d1-dist >= 100 && dist <= 20 {
					count++
				}
			}
		}
		return count
	}

	shortcuts := findShortcuts(path)
	stepsSaved := make([]int, 100)
	for _, shortcut := range shortcuts {
		saved := shortcut.stepsSaved
		if saved > 99 {
			saved = 99
		}
		stepsSaved[saved]++
	}

	last := 0
	for _, saved := range stepsSaved {
		if saved > 0 {
			// fmt.Printf("There are %d cheats that save %d picoseconds.\n", saved, i)
			last = saved
		}
	}
	return last
}

type Position struct {
	x, y int
}

type Step struct {
	pos   Position
	score int
}

func parseInput(input string) ([][]string, Position, Position) {
	grid := make([][]string, 0)
	var start, end Position
	for y, line := range strings.Split(input, "\n") {
		newRow := make([]string, 0)
		for x, cell := range line {
			newRow = append(newRow, string(cell))

			switch cell {
			case 'S':
				start = Position{x, y}
			case 'E':
				end = Position{x, y}
			}
		}
		grid = append(grid, newRow)
	}

	return grid, start, end
}

func dijkstras(grid [][]string, start Position) (int, []Position) {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})

	priorityQueue.Enqueue(Step{start, 0})
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

		if grid[y][x] == "E" {
			path := make([]Position, 0)
			pos := current.pos
			for pos.x != 0 && pos.y != 0 {
				path = append(path, pos)
				pos = predecessorMap[pos]
			}
			slices.Reverse(path)
			return current.score, path
		}

		for _, dir := range directions {
			newPos := Position{current.pos.x + dir.x, current.pos.y + dir.y}
			if isInGrid(newPos, grid) && !visited[newPos.y][newPos.x] && grid[newPos.y][newPos.x] != "#" {
				predecessorMap[newPos] = current.pos
				priorityQueue.Enqueue(Step{newPos, current.score + 1})
			}
		}
	}

	return -1, []Position{}
}

func isInGrid(pos Position, grid [][]string) bool {
	return pos.y >= 0 && pos.y < len(grid) && pos.x >= 0 && pos.x < len(grid[0])
}

type Shortcut struct {
	start       Position
	end         Position
	shortcutPos Position
	stepsSaved  int
}

func findShortcuts(path []Position) []Shortcut {
	shortcuts := []Shortcut{}
	pathMap := make(map[Position]struct{}, len(path))
	for _, pos := range path {
		pathMap[pos] = struct{}{}
	}

	for i := 0; i < len(path)-2; i++ {
		for j := i + 2; j < len(path); j++ {
			if canCreateShortcut(path[i], path[j]) {
				shortcutPoint := getShortcutPoint(path[i], path[j])

				// it's only a shortcut if it's not already in the path
				if _, exists := pathMap[shortcutPoint]; !exists {
					stepsSaved := j - i - 2
					shortcuts = append(shortcuts, Shortcut{
						start:       path[i],
						end:         path[j],
						shortcutPos: shortcutPoint,
						stepsSaved:  stepsSaved,
					})
				}
			}
		}
	}

	return shortcuts
}

func canCreateShortcut(p1, p2 Position) bool {
	dx := abs(p2.x - p1.x)
	dy := abs(p2.y - p1.y)
	return dx+dy == 2
}

func getShortcutPoint(p1, p2 Position) Position {
	if p1.x == p2.x {
		return Position{p1.x, (p1.y + p2.y) / 2}
	}
	return Position{(p1.x + p2.x) / 2, p1.y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
