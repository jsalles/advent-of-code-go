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
	grid := ParseInput(input)
	seen := make(map[Point]bool)
	var areas []Area

	for row, line := range grid {
		for col := range line {
			if seen[Point{row, col}] {
				continue
			}

			newArea := bfs(grid, row, col, &seen)
			areas = append(areas, newArea)
		}
	}

	sum := 0
	for _, area := range areas {
		if part2 {
			sum += len(area.positions) * CalculateSides(area.positions)
		} else {
			sum += len(area.positions) * CalculatePerimeter(area.positions)
		}
	}

	return sum
}

var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func bfs(grid [][]string, row int, col int, seen *map[Point]bool) Area {
	area := Area{grid[row][col], make([]Point, 0)}

	queue := list.New()
	queue.PushBack(Point{row, col})
	area.positions = append(area.positions, Point{row, col})
	(*seen)[Point{row, col}] = true

	for queue.Len() > 0 {
		queueElement := queue.Back()
		queue.Remove(queueElement)
		current := queueElement.Value.(Point)

		for _, direction := range directions {
			newPos := Point{current.y + direction.y, current.x + direction.x}
			if newPos.y < 0 || newPos.y >= len(grid) || newPos.x < 0 || newPos.x >= len(grid[0]) {
				continue
			}
			if !(*seen)[newPos] && grid[newPos.y][newPos.x] == area.value {
				queue.PushBack(newPos)
				area.positions = append(area.positions, newPos)
				(*seen)[newPos] = true
			}
		}
	}

	return area
}

func CalculatePerimeter(positions []Point) int {
	positonMap := make(map[Point]bool)
	for _, pos := range positions {
		positonMap[pos] = true
	}

	perimeter := 0
	for _, pos := range positions {
		neighbors := []Point{
			{pos.y - 1, pos.x},
			{pos.y, pos.x + 1},
			{pos.y + 1, pos.x},
			{pos.y, pos.x - 1},
		}

		for _, neighbor := range neighbors {
			if !positonMap[neighbor] {
				perimeter++
			}
		}
	}

	return perimeter
}

type Normal struct {
	pos Point
	dir Point
}

func CalculateSides(positions []Point) int {
	positionsMap := make(map[Point]bool)
	for _, pos := range positions {
		positionsMap[pos] = true
	}

	normals := make(map[Normal]bool)
	for pos := range positionsMap {
		for _, dir := range directions {
			neighbor := Point{pos.y + dir.y, pos.x + dir.x}
			if !positionsMap[neighbor] {
				normals[Normal{pos, dir}] = true
			}
		}
	}

	finalNormals := make(map[Normal]bool)

	for normal := range normals {
		// the rotation magic helps us find the duplicate normals in our list
		rotated := Point{-normal.dir.x, normal.dir.y}
		neighborPos := Point{normal.pos.y + rotated.y, normal.pos.x + rotated.x}
		neighborNormal := Normal{neighborPos, normal.dir}

		if !normals[neighborNormal] {
			finalNormals[normal] = true
		}
	}

	return len(finalNormals)
}

type Point struct {
	y, x int
}

type Area struct {
	value     string
	positions []Point
}

func ParseInput(input string) [][]string {
	var grid [][]string

	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		var newLine []string
		for _, cell := range line {
			newLine = append(newLine, string(cell))
		}
		grid = append(grid, newLine)
	}

	return grid
}
