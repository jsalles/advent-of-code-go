package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

const STEPS = 100

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	robots, dims := ParseInput(input)
	if part2 {
		// 8822, 8923 9024
		k := 0
		for range 10000 {
			for i := 0; i < len(robots); i++ {
				robots[i] = MoveRobot(robots[i], dims, 1)
			}
			k++
			// if (k+66)%101 == 0 {
			// 	fmt.Println(k)
			// 	PrintBoard(dims, robots, false)
			// 	fmt.Println()
			// }

		}
		return 0
	}

	// PrintBoard(dims, robots, false)
	quadrantCount := make([]int, 4)
	for i := 0; i < len(robots); i++ {
		robot := MoveRobot(robots[i], dims, STEPS)
		robots[i] = robot
		if valid, quadrant := FindQuadrant(robot.pos, dims); valid {
			quadrantCount[quadrant]++
		}
	}

	sum := quadrantCount[0]
	for i := 1; i < len(quadrantCount); i++ {
		sum *= quadrantCount[i]
	}
	return sum
}

type Point struct {
	x, y int
}

type Robot struct {
	pos, vel Point
}

func ParseInput(input string) ([]Robot, Point) {
	var robots []Robot
	re := regexp.MustCompile(`p=(\d+),(\d+) v=([-\d]+),([-\d]+)$`)
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	dims := strings.Split(lines[0], " ")
	width, _ := strconv.Atoi(dims[0])
	height, _ := strconv.Atoi(dims[1])
	for _, line := range lines[1:] {
		matches := re.FindStringSubmatch(line)
		var robot Robot
		robot.pos.x, _ = strconv.Atoi(matches[1])
		robot.pos.y, _ = strconv.Atoi(matches[2])
		robot.vel.x, _ = strconv.Atoi(matches[3])
		robot.vel.y, _ = strconv.Atoi(matches[4])

		robots = append(robots, robot)

	}

	return robots, Point{width, height}
}

func MoveRobot(robot Robot, dims Point, steps int) Robot {
	robot.pos.x = (robot.pos.x + robot.vel.x*steps) % dims.x
	if robot.pos.x < 0 {
		robot.pos.x += dims.x
	}
	robot.pos.y = (robot.pos.y + robot.vel.y*steps) % dims.y
	if robot.pos.y < 0 {
		robot.pos.y += dims.y
	}
	return robot
}

func FindQuadrant(pos, dims Point) (bool, int) {
	midX := int(math.Floor(float64(dims.x) / 2))
	midY := int(math.Floor(float64(dims.y) / 2))
	if pos.x == midX || pos.y == midY {
		return false, -1
	}
	if pos.x < midX {
		if pos.y < midY {
			return true, 0
		} else {
			return true, 2
		}
	} else {
		if pos.y < midY {
			return true, 1
		} else {
			return true, 3
		}
	}
}

func PrintBoard(dims Point, robots []Robot, skipMiddle bool) {
	midX := int(math.Floor(float64(dims.x) / 2))
	midY := int(math.Floor(float64(dims.y) / 2))
	for y := 0; y < dims.y; y++ {
		if skipMiddle && y == midY {
			fmt.Println()
			continue
		}
		for x := 0; x < dims.x; x++ {
			if skipMiddle && x == midX {
				fmt.Print(" ")
				continue
			}
			robotCount := 0
			for _, robot := range robots {
				if robot.pos.x == x && robot.pos.y == y {
					robotCount++
				}
			}
			if robotCount == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(robotCount)
			}
		}
		fmt.Println()
	}
}
