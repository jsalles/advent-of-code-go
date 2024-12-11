package main

import (
	"fmt"
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
	blinkCount := 25
	if part2 {
		blinkCount = 75
	}
	nums := ParseInput(input)
	sum := 0
	for _, num := range nums {
		sum += CountStonesAfterBlink(num, blinkCount)
	}

	return sum
}

func PrintDp() {
	for num, counts := range dp {
		fmt.Print(num, " => ")
		for i, count := range counts {
			if count != -1 {
				fmt.Print(i, ": ", count)
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

func ParseInput(input string) []int {
	splitInput := strings.Split(strings.Trim(input, "\n"), " ")
	nums := make([]int, len(splitInput))
	for i := 0; i < len(splitInput); i++ {
		nums[i], _ = strconv.Atoi(splitInput[i])
	}

	return nums
}

var dp = make(map[int][]int) // map[num][blinkCount] = howManyStonesAfterBlinks

func CountStonesAfterBlink(num int, remaining int) int {
	if _, exists := dp[num]; exists {
		if dp[num][remaining-1] != -1 {
			return dp[num][remaining-1]
		}
	} else {
		dp[num] = make([]int, 75)
		for i := 0; i < 75; i++ {
			dp[num][i] = -1
		}
	}

	if remaining == 1 {
		result := Blink(num)
		dp[num][remaining-1] = len(result)
		return len(result)
	}

	sum := 0
	for _, num := range Blink(num) {
		sum += CountStonesAfterBlink(num, remaining-1)
	}

	dp[num][remaining-1] = sum
	return sum
}

func Blink(num int) []int {
	newNums := make([]int, 0)

	if num == 0 {
		newNums = append(newNums, 1)
	} else if digits := countDigits(num); digits%2 == 0 {
		left, right := splitNumber(num)
		newNums = append(newNums, left, right)
	} else {
		newNums = append(newNums, num*2024)
	}

	return newNums
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func splitNumber(n int) (int, int) {
	digits := countDigits(n)
	divisor := 1
	for i := 0; i < digits/2; i++ {
		divisor *= 10
	}
	return n / divisor, n % divisor
}
