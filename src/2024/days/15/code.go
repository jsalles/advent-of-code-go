package main

import (
	"bufio"
	"fmt"
	"os"
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
	grid, dims, robotPos, commands := ParseInput(input)
	if part2 {
		grid, robotPos = ExpandGrid(grid)
		dims.width *= 2
	}

	grid = runCommands(grid, dims, robotPos, commands)

	sum := 0
	for row, line := range grid {
		for col, cell := range line {
			if cell == "O" || cell == "[" {
				sum += 100*row + col
			}
		}
	}

	return sum
}

type GridDims struct {
	width, height int
}

type Position struct {
	x, y int
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func GetDirection(strDirection string) Position {
	switch strDirection {
	case "^":
		return Position{0, -1}
	case ">":
		return Position{1, 0}
	case "v":
		return Position{0, 1}
	case "<":
		return Position{-1, 0}
	}

	panic(fmt.Sprintf("'%s' not supported", strDirection))
}

func WaitForEnter() {
	reader := bufio.NewReader(os.Stdin)

	// Read a single character
	_, _, err := reader.ReadRune()
	if err != nil {
		return
	}
}

func runCommands(grid [][]string, dims GridDims, robotPos Position, commands []string) [][]string {
	for _, command := range commands {
		dir := GetDirection(command)
		newCell := grid[robotPos.y+dir.y][robotPos.x+dir.x]
		if newCell == "[" || newCell == "]" && dir.y != 0 {
			grid, robotPos = PushDoubleBoxesVertically(grid, dims, robotPos, dir)
		} else {
			grid, robotPos = PushSimpleBoxes(grid, dims, robotPos, dir)
		}

	}
	return grid
}

func PushDoubleBoxesVertically(grid [][]string, dims GridDims, robotPos Position, dir Position) ([][]string, Position) {
	newPos := Position{robotPos.x + dir.x, robotPos.y + dir.y}
	queue := []Position{newPos}
	if grid[newPos.y][newPos.x] == "]" {
		queue = append(queue, Position{newPos.x - 1, newPos.y})
	} else {
		queue = append(queue, Position{newPos.x + 1, newPos.y})
	}

	visited := make(map[Position]struct{})
	visitedSlice := []Position{}

	for len(queue) != 0 {
		currentPos := queue[0]
		queue = queue[1:]

		if _, exists := visited[currentPos]; exists {
			continue
		}

		visited[currentPos] = struct{}{}
		visitedSlice = append(visitedSlice, currentPos)

		newX, newY := currentPos.x+dir.x, currentPos.y+dir.y
		switch grid[newY][newX] {
		case ".":
			continue
		case "#":
			return grid, robotPos
		case "]":
			queue = append(queue, Position{newX, newY})
			queue = append(queue, Position{newX - 1, newY})
		case "[":
			queue = append(queue, Position{newX, newY})
			queue = append(queue, Position{newX + 1, newY})
		}
	}

	// move all the cells visited by the bfs
	for i := len(visitedSlice) - 1; i >= 0; i-- {
		x, y := visitedSlice[i].x+dir.x, visitedSlice[i].y+dir.y
		grid[y][x] = grid[visitedSlice[i].y][visitedSlice[i].x]
		grid[visitedSlice[i].y][visitedSlice[i].x] = "."
	}

	// also shift the robot
	grid[robotPos.y][robotPos.x] = "."
	grid[robotPos.y+dir.y][robotPos.x+dir.x] = "@"
	return grid, Position{robotPos.x + dir.x, robotPos.y + dir.y}
}

func PushSimpleBoxes(grid [][]string, dims GridDims, robotPos Position, dir Position) ([][]string, Position) {
	valid, emptyPos := CanMoveRobotHorizontally(grid, robotPos, dims, dir)
	// can't move
	if !valid {
		return grid, robotPos
	}
	for pos := emptyPos; pos != robotPos; {
		neighbor := Position{pos.x - dir.x, pos.y - dir.y}
		grid[pos.y][pos.x], grid[neighbor.y][neighbor.x] = grid[neighbor.y][neighbor.x], grid[pos.y][pos.x]
		pos = neighbor
	}

	robotPos.x += dir.x
	robotPos.y += dir.y
	return grid, robotPos
}

func IsInGrid(pos Position, dims GridDims) bool {
	return pos.x >= 0 && pos.x < dims.width && pos.y >= 0 && pos.y < dims.height
}

func CanMoveRobotHorizontally(grid [][]string, pos Position, dims GridDims, direction Position) (bool, Position) {
	pos.x += direction.x
	pos.y += direction.y
	for IsInGrid(pos, dims) && grid[pos.y][pos.x] != "." && grid[pos.y][pos.x] != "#" {
		pos.x += direction.x
		pos.y += direction.y
	}

	return IsInGrid(pos, dims) && grid[pos.y][pos.x] == ".", pos
}

func ParseInput(input string) ([][]string, GridDims, Position, []string,
) {
	blocks := strings.Split(input, "\n\n")
	gridWidth := strings.Index(blocks[0], "\n")
	gridHeight := strings.Count(blocks[0], "\n") + 1
	grid := make([][]string, gridHeight)
	for row := range gridHeight {
		grid[row] = make([]string, gridWidth)
	}

	var robotStart Position
	for row, line := range strings.Split(blocks[0], "\n") {
		for col, cell := range line {
			grid[row][col] = string(cell)

			if grid[row][col] == "@" {
				robotStart = Position{col, row}
			}
		}
	}

	var commands []string
	for _, line := range strings.Split(blocks[1], "\n") {
		for _, command := range line {
			commands = append(commands, string(command))
		}
	}

	return grid, GridDims{gridWidth, gridHeight}, robotStart, commands
}

func ExpandGrid(grid [][]string) ([][]string, Position) {
	var newGrid [][]string
	var robotPos Position
	for y, line := range grid {
		var newLine []string
		for x, cell := range line {
			switch cell {
			case "#":
				newLine = append(newLine, "#", "#")
			case ".":
				newLine = append(newLine, ".", ".")
			case "O":
				newLine = append(newLine, "[", "]")
			case "@":
				{
					newLine = append(newLine, "@", ".")
					robotPos = Position{2 * x, y}
				}
			}
		}
		newGrid = append(newGrid, newLine)
	}

	return newGrid, robotPos
}
