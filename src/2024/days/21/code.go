package main

import (
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
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	numDirectionals := 2
	if part2 {
		numDirectionals = 25
	}
	// solve part 1 here
	sum := 0
	mem := make(map[string][]int)
	for _, code := range strings.Split(input, "\n") {
		start := keypad['A']
		keypadSequence := typeInKeypad(start, code)
		directionalLength := typeInChainDirectionals(keypadSequence, numDirectionals, 1, mem)

		numeric := extractNumeric(code)
		sum += numeric * directionalLength
	}

	return sum
}

type Position struct {
	x, y int
}

var keypad = map[rune]Position{
	'7': {0, 3},
	'8': {1, 3},
	'9': {2, 3},
	'4': {0, 2},
	'5': {1, 2},
	'6': {2, 2},
	'1': {0, 1},
	'2': {1, 1},
	'3': {2, 1},
	'0': {1, 0},
	'A': {2, 0},
}

func typeInChainDirectionals(input string, numDirectionals int, curDirectional int, mem map[string][]int) int {
	if val, exists := mem[input]; exists {
		if val[curDirectional-1] != 0 {
			return val[curDirectional-1]
		}
	} else {
		mem[input] = make([]int, numDirectionals)
	}

	directionalSequence := typeInDirectionalPad(input, directional, directional['A'])
	mem[input][0] = len(directionalSequence)

	if curDirectional == numDirectionals {
		return len(directionalSequence)
	}

	splitSeq := make([]string, 0)
	for _, seq := range strings.Split(strings.Trim(directionalSequence, "A"), "A") {
		splitSeq = append(splitSeq, seq+"A")
	}

	count := 0
	for _, step := range splitSeq {
		c := typeInChainDirectionals(step, numDirectionals, curDirectional+1, mem)
		if _, ok := mem[step]; !ok {
			mem[step] = make([]int, numDirectionals)
		}
		mem[step][0] = c
		count += c
	}

	mem[input][curDirectional-1] = count
	return count
}

func typeInKeypad(pos Position, code string) string {
	output := ""
	for _, bit := range code {
		dest := keypad[bit]
		dx := dest.x - pos.x
		dy := dest.y - pos.y

		vertical := ""
		horizontal := ""
		// going down
		for dy < 0 {
			vertical += "v"
			dy++
		}
		// going up
		for dy > 0 {
			vertical += "^"
			dy--
		}
		// going left
		for dx < 0 {
			horizontal += "<"
			dx++
		}
		// going right
		for dx > 0 {
			horizontal += ">"
			dx--
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >
		if pos.y == 0 && dest.x == 0 {
			output += vertical + horizontal
		} else if pos.x == 0 && dest.y == 0 {
			output += horizontal + vertical
		} else if dest.x-pos.x < 0 {
			output += horizontal + vertical
		} else if dest.x-pos.x >= 0 {
			output += vertical + horizontal
		}

		output += "A"
		pos = keypad[bit]
	}
	return output
}

var directional = map[rune]Position{
	'^': {1, 1},
	'A': {2, 1},
	'<': {0, 0},
	'v': {1, 0},
	'>': {2, 0},
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func typeInDirectionalPad(input string, keypad map[rune]Position, start Position) string {
	current := start
	output := ""

	for _, char := range input {
		dest := keypad[char]
		diffX, diffY := dest.x-current.x, dest.y-current.y

		horizontal, vertical := "", ""

		for i := 0; i < Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal += ">"
			} else {
				horizontal += "<"
			}
		}

		for i := 0; i < Abs(diffY); i++ {
			if diffY >= 0 {
				vertical += "^"
			} else {
				vertical += "v"
			}
		}

		// prioritisation order:
		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if current.x == 0 && dest.y == 1 {
			output += horizontal + vertical
		} else if current.y == 1 && dest.x == 0 {
			horizontal += vertical + horizontal
		} else if diffX < 0 {
			output += horizontal + vertical
		} else if diffX >= 0 {
			output += vertical + horizontal
		}
		current = dest
		output += "A"
	}
	return output
}

func typeInDirectional(pos Position, code string) string {
	output := ""
	for _, bit := range code {
		dest := directional[bit]
		dx := dest.x - pos.x
		dy := dest.y - pos.y

		horizontal := ""
		vertical := ""
		// going left
		for dx < 0 {
			horizontal += "<"
			dx++
		}
		// going up
		for dy > 0 {
			vertical += "^"
			dy--
		}
		// going right
		for dx > 0 {
			horizontal += ">"
			dx--
		}
		// going down
		for dy < 0 {
			vertical += "v"
			dy++
		}

		if pos.x == 0 && dest.y == 1 {
			output += horizontal + vertical
		} else if pos.y == 1 && dest.x == 0 {
			output += vertical + horizontal
		} else if dest.x-pos.x < 0 {
			output += horizontal + vertical
		} else if dest.x-pos.x >= 0 {
			output += vertical + horizontal
		}

		output += "A"
		pos = directional[bit]
	}
	return output
}

func extractNumeric(str string) int {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindString(str)
	num, _ := strconv.Atoi(numbers)
	return num
}
