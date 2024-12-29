package main

import (
	"strconv"
	"strings"
	"unicode"

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
	ans1, ans2 := calculateSecretNums(strings.Split(input, "\n"))
	if part2 {
		return ans2
	}
	return ans1
}

func findSecretNumber(input int64) int64 {
	input = prune(mix(int64(64)*input, input))
	input = prune(mix(input/32, input))
	input = prune(mix(input*2048, input))
	return input
}

func mix(value, secretNum int64) int64 {
	return value ^ secretNum
}

func prune(value int64) int64 {
	return value % 16777216
}

type bananaPrice struct {
	num    int
	change int
}

func getNumAndChange(input int64, previous int) bananaPrice {
	return bananaPrice{num: int(input % 10), change: int(input%10) - previous}
}

func FetchNumFromStringIgnoringNonNumeric(line string) int {
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(localNum)
	}
	return 0
}

func calculateSecretNums(input []string) (int64, int) {
	var sum int64
	b := make([][]bananaPrice, len(input))
	for i, line := range input {
		b[i] = make([]bananaPrice, 2000)
		num := int64(FetchNumFromStringIgnoringNonNumeric(line))
		prev := int(num % 10)
		for j := range 2000 {
			num = findSecretNumber(num)
			b[i][j] = getNumAndChange(num, prev)
			prev = int(num % 10)
		}
		sum += num
	}

	return sum, findMaxNumberOfBananas(b)
}

type seq struct {
	a, b, c, d int
}

func findMaxNumberOfBananas(b [][]bananaPrice) int {
	seqMap := make(map[seq][]int)
	for i := range b {
		s := []int{b[i][0].change, b[i][1].change, b[i][2].change}
		for j := 3; j < len(b[i]); j++ {
			s = append(s, b[i][j].change)

			if _, ok := seqMap[seq{s[0], s[1], s[2], s[3]}]; !ok {
				seqMap[seq{s[0], s[1], s[2], s[3]}] = make([]int, len(b))
			}

			if seqMap[seq{s[0], s[1], s[2], s[3]}][i] == 0 {
				seqMap[seq{s[0], s[1], s[2], s[3]}][i] = b[i][j].num
			}

			s = s[1:]
		}
	}
	max := 0
	for _, n := range seqMap {
		sum := 0
		for _, r := range n {
			sum += r
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
