package main

import (
	"fmt"
	"slices"
	"sort"
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
	nodes, instructions := parseInput(input)
	changed := true
	for changed {
		changed = false
		for _, instr := range instructions {
			previous := nodes[instr.target]
			nodes[instr.target] = performOperation(nodes, instr.firstInput, instr.secondInput, instr.operator)
			if nodes[instr.target] != previous {
				changed = true
			}
		}
	}
	if part2 {
		xNodes, yNodes, zNodes := captureNodes(nodes, 'x'), captureNodes(nodes, 'y'), captureNodes(nodes, 'z')
		flipped := addBits(xNodes, yNodes, zNodes)

		sort.Strings(flipped)
		return strings.Join(flipped, ",")
	}

	bits := captureNodes(nodes, 'z')
	return convertToInt(bits)
}

func printBits(xNodes, yNodes, zNodes []int) {
	for i := 45; i >= 0; i-- {
		fmt.Print(i / 10)
	}
	fmt.Println()
	for i := 45; i >= 0; i-- {
		fmt.Print(i % 10)
	}
	fmt.Println()
	fmt.Println()

	fmt.Printf(" %s\n", convertToString(xNodes))
	fmt.Printf(" %s\n", convertToString(yNodes))
	fmt.Println()
	fmt.Println(convertToString(zNodes))
}

func addBits(first, second, check []int) []string {
	var flipped []string

	slices.Reverse(first)
	slices.Reverse(second)
	slices.Reverse(check)
	// 18 is flipped: switch hmt, z18
	flipped = append(flipped, "hmt", "z18")
	// 27 is flipped: switch z27, bfq
	flipped = append(flipped, "z27", "bfq")
	// 31 is flipped: switch z31, hkh
	flipped = append(flipped, "z31", "hkh")
	// 39 is flipped: switch bng, fjp
	flipped = append(flipped, "bng", "fjp")
	carry := 0
	for i := 0; i < min(45, len(first)); i++ {
		result := (first[i] + second[i] + carry) % 2
		if result != check[i] {
			// fmt.Printf("z%02d is flipped\n", i)
		}
		carry = (first[i] + second[i] + carry) / 2
	}

	if carry != check[len(check)-1] {
		// fmt.Println("last")
	}

	return flipped
}

func convertToString(bits []int) string {
	str := make([]string, len(bits))
	for i, bit := range bits {
		str[i] = strconv.Itoa(bit)
	}
	return strings.Join(str, "")
}

func captureNodes(nodes map[string]bool, initial byte) []int {
	capturedNodes := make([]string, 0)
	for node := range nodes {
		if node[0] == initial {
			capturedNodes = append(capturedNodes, node)
		}
	}
	slices.SortFunc(capturedNodes, func(a, b string) int {
		if a < b {
			return 1
		}
		return -1
	})

	bits := make([]int, len(capturedNodes))
	for i := 0; i < len(capturedNodes); i++ {
		if nodes[capturedNodes[i]] {
			bits[i] = 1
		} else {
			bits[i] = 0
		}
	}

	return bits
}

func performOperation(nodes map[string]bool, first, second, operator string) bool {
	switch operator {
	case "AND":
		return nodes[first] && nodes[second]
	case "OR":
		return nodes[first] || nodes[second]
	case "XOR":
		return nodes[first] != nodes[second]
	}

	panic("unexpected operator")
}

func convertToInt(bits []int) int64 {
	result := int64(0)
	for i := 0; i < len(bits); i++ {
		if bits[len(bits)-i-1] == 1 {
			result |= 1 << i
		}
	}
	return result
}

type Instruction struct {
	firstInput, secondInput string
	target                  string
	operator                string
}

func parseInput(input string) (map[string]bool, []Instruction) {
	blocks := strings.Split(input, "\n\n")

	nodes := make(map[string]bool)
	for _, line := range strings.Split(blocks[0], "\n") {
		node := line[:3]
		val := false
		if line[5:] == "1" {
			val = true
		}

		nodes[node] = val
	}

	instructions := make([]Instruction, 0)
	for _, line := range strings.Split(strings.Trim(blocks[1], "\n"), "\n") {
		parts := strings.Split(line, " ")
		var instr Instruction
		instr.firstInput = parts[0]
		instr.operator = parts[1]
		instr.secondInput = parts[2]
		instr.target = parts[4]
		instructions = append(instructions, instr)
	}

	return nodes, instructions
}
