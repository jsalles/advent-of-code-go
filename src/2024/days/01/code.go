package main

import (
	pqueue "aoc-in-go/src/2024/utils/heap_utils"
	"math"
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
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part_2(input)
	}
	// solve part 1 here
	pq_left, pq_right := parse_input(input)

	distance := 0
	for pq_left.Len() > 0 && pq_right.Len() > 0 {
		left := pq_left.Pop().Value
		right := pq_right.Pop().Value

		distance = distance + int(math.Abs(float64(right-left)))
	}

	return distance
}

func parse_input(input string) (*pqueue.PriorityQueue[int], *pqueue.PriorityQueue[int]) {
	lines := strings.Split(input, "\n")
	pq_left := pqueue.New[int]()
	pq_right := pqueue.New[int]()

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			left_num, _ := strconv.Atoi(numbers[0])
			pq_left.Push(left_num, left_num)

			right_num, _ := strconv.Atoi(numbers[1])
			pq_right.Push(right_num, right_num)
		}
	}

	return pq_left, pq_right
}

func part_2(input string) any {
	var similarity int
	pq_left, pq_right := parse_input(input)
	for !pq_left.IsEmpty() && !pq_right.IsEmpty() {
		for !pq_right.IsEmpty() && pq_right.Peek().Value < pq_left.Peek().Value {
			pq_right.Pop()
		}
		if pq_right.IsEmpty() {
			break
		}

		for !pq_left.IsEmpty() && pq_left.Peek().Value < pq_right.Peek().Value {
			pq_left.Pop()
		}
		if pq_left.IsEmpty() {
			break
		}

		count := 0
		current_left := pq_left.Peek().Value
		for current_left == pq_right.Peek().Value && !pq_right.IsEmpty() {
			pq_right.Pop()
			count++
		}
		for count > 0 && !pq_left.IsEmpty() && pq_left.Peek().Value == current_left {
			pq_left.Pop()
			similarity += current_left * count
		}

	}

	return similarity
}
