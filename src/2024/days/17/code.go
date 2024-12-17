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
	state, commands := parseInput(input)

	if part2 {
		amountForCopy := findDigitalCopy(&state, commands)
		return amountForCopy

	}

	result := processOpCodes(&state, commands)
	output := make([]string, len(result))
	for i, op := range result {
		output[i] = strconv.Itoa(op)
	}
	return strings.Join(output, ",")
}

type State struct {
	A, B, C  uint64 // registers
	instrPtr int    // instruction pointer
}

func parseInput(input string) (State, []int) {
	var state State
	var program []int

	lines := strings.Split(input, "\n")

	fmt.Sscanf(lines[0], "Register A: %d", &state.A)
	fmt.Sscanf(lines[1], "Register B: %d", &state.B)
	fmt.Sscanf(lines[2], "Register C: %d", &state.C)

	programStr := strings.TrimPrefix(lines[4], "Program: ")
	for _, instruction := range strings.Split(programStr, ",") {
		if num, err := strconv.Atoi(strings.TrimSpace(instruction)); err == nil {
			program = append(program, num)
		}
	}

	return state, program
}

// From reddit: Notice that each digit is dependent on the rightmost 10 bits of A, then A shifts 3 places right until 0.
// So generate all possible combinations of 7bit numbers (1-128) than proceed to guess leftmost 3bits as in DFS while checking the result.
//
// Very nice catch!
func dfs(state *State, program []int, expected []int, acc uint64, level int) uint64 {
	if len(expected) == 0 {
		return acc
	}

	for k := uint64(0); k < 8; k++ {
		testValue := (k << ((3 * level) + 7)) | acc

		state.A = testValue >> (3 * level)
		state.B = 0
		state.C = 0
		result := processOpCodes(state, program)
		if result[0] == expected[0] {
			if found := dfs(state, program, expected[1:], testValue, level+1); found > 0 {
				return found
			}
		}
	}
	return 0
}

func findDigitalCopy(state *State, program []int) uint64 {
	var currentVal uint64
	for testVal := uint64(0); testVal < 128; testVal++ {
		if result := dfs(state, program, program, testVal, 0); result > 0 {
			currentVal = result
			break
		}
	}
	return currentVal
}

func getOperandValue(state *State, operand int) uint64 {
	switch operand {
	case 4:
		return state.A
	case 5:
		return state.B
	case 6:
		return state.C
	default:
		return uint64(operand)
	}
}

func processOpCodes(state *State, commands []int) []int {
	state.instrPtr = 0
	output := []int{}

	for state.instrPtr < len(commands) {
		opCode := commands[state.instrPtr]
		arg := commands[state.instrPtr+1]
		state.instrPtr += 2

		switch opCode {
		case 0: // Advance
			state.A >>= getOperandValue(state, arg)
		case 1: // BitXorLiteral
			state.B ^= uint64(arg)
		case 2: // BitStore
			state.B = getOperandValue(state, arg) & 7
		case 3: // JumpNotZero
			if state.A != 0 {
				state.instrPtr = arg
			}
		case 4: // BitXorC
			state.B ^= state.C
		case 5: // Output
			output = append(output, int(getOperandValue(state, arg)&7))
		case 6: // BitDivideA
			state.B = state.A >> getOperandValue(state, arg)
		case 7: // CopyDivideA
			state.C = state.A >> getOperandValue(state, arg)
		}
	}
	return output
}
