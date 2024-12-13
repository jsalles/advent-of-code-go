package main

import (
	"math"
	"regexp"
	"strconv"
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

const PART_2_ERROR = 10000000000000

func run(part2 bool, input string) any {
	games := ParseInput(input)
	sum := 0
	for _, game := range games {
		prize := game.prize
		if part2 {
			prize.x += PART_2_ERROR
			prize.y += PART_2_ERROR
		}
		sum += countPressesEEA(game.a, game.b, prize)
	}
	return sum
}

// Part2 pushed the numbers too far. Extended Euclidean Algorithm solves it in O(1). TIL
func countPressesEEA(a, b Button, prize Point) int {
	// First equation coefficients
	c1 := a.shift.x
	c2 := b.shift.x
	k1 := prize.x

	// Second equation coefficients
	c3 := a.shift.y
	c4 := b.shift.y
	k2 := prize.y

	// Solve using determinants
	det := c1*c4 - c2*c3
	if det == 0 {
		return 0 // no solution
	}

	// Find countA and countB using Cramer's rule
	countA := (k1*c4 - k2*c2) / det
	countB := (c1*k2 - c3*k1) / det

	// If countA and countB are not integers or are negative, no solution
	if countA < 0 || countB < 0 || float64(countA) != float64((k1*c4-k2*c2))/float64(det) ||
		float64(countB) != float64((c1*k2-c3*k1))/float64(det) {
		return 0
	}

	return countA*a.cost + countB*b.cost
}

// countA * a.shift.X + countB * b.shift.X = prize.X
// countA * a.shift.Y + countB * b.shift.Y = prize.Y
// minimize for 3 * countA + countB
func countPresses(a, b Button, prize Point) int {
	// countA(a.shift.x) ≤ prize.x
	maxCountA := prize.x/a.shift.x + 1

	minCost := math.MaxInt32

	for countA := 0; countA <= maxCountA; countA++ {
		// prize.x = countA(a.shift.x) + countB(b.shift.x)
		// countB = (prize.x - x(a.shift.x)) / b.shift.x
		if (prize.x-countA*a.shift.x)%b.shift.x != 0 {
			continue // countB must be int
		}
		countB := (prize.x - countA*a.shift.x) / b.shift.x
		if countB < 0 {
			continue
		}

		if countA*a.shift.y+countB*b.shift.y == prize.y {
			cost := countA*a.cost + countB*b.cost
			if cost < minCost {
				minCost = cost
			}
		}
	}

	if minCost == math.MaxInt32 {
		return 0 // no solution found
	}
	return minCost
}

type Point struct {
	x, y int
}

type Button struct {
	cost  int
	shift Point
}

type Game struct {
	a, b  Button
	prize Point
}

func ParseInput(input string) []Game {
	var games []Game

	reButton := regexp.MustCompile(`^Button \w: X\+(\d+), Y\+(\d+)$`)
	rePrize := regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)
	for _, gameDetails := range strings.Split(input, "\n\n") {
		var game Game
		lines := strings.Split(gameDetails, "\n")

		// Button A: X+94, Y+34
		matches := reButton.FindStringSubmatch(lines[0])
		game.a.shift.x, _ = strconv.Atoi(matches[1])
		game.a.shift.y, _ = strconv.Atoi(matches[2])
		game.a.cost = 3

		matches = reButton.FindStringSubmatch(lines[1])
		game.b.shift.x, _ = strconv.Atoi(matches[1])
		game.b.shift.y, _ = strconv.Atoi(matches[2])
		game.b.cost = 1

		// Prize: X=8400, Y=5400
		matches = rePrize.FindStringSubmatch(lines[2])
		game.prize.x, _ = strconv.Atoi(matches[1])
		game.prize.y, _ = strconv.Atoi(matches[2])

		games = append(games, game)
	}

	return games
}
