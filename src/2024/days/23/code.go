package main

import (
	"sort"
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
	graph := parseInput(input)
	if part2 {
		return findMaxCliques(graph)
	}
	// solve part 1 here
	triples := findConnectedTriples(graph)
	count := 0
	for triple := range triples {
		if triple.first[0] == 't' || triple.second[0] == 't' || triple.third[0] == 't' {
			count++
		}
	}
	return count
}

func parseInput(input string) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		nodes := strings.Split(line, "-")
		first, second := nodes[0], nodes[1]
		if _, exists := graph[first]; !exists {
			graph[first] = make(map[string]bool)
		}
		if _, exists := graph[second]; !exists {
			graph[second] = make(map[string]bool)
		}

		graph[first][second] = true
		graph[second][first] = true
	}

	return graph
}

type Triple struct {
	first, second, third string
}

func findConnectedTriples(graph map[string]map[string]bool) map[Triple]struct{} {
	triangles := make(map[Triple]struct{})

	for first, neighbors := range graph {
		for second := range neighbors {
			if first < second {
				for third := range neighbors {
					if second < third {
						if graph[second][third] {
							triangles[Triple{first, second, third}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return triangles
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func BronKerbosch(currentClique []string, yetToConsider []string, alreadyConsidered []string, lanMap map[string]map[string]bool, cliques [][]string) [][]string {
	if len(yetToConsider) == 0 && len(alreadyConsidered) == 0 {
		cliques = append(cliques, append([]string{}, currentClique...))
		return cliques
	}

	for index := 0; index < len(yetToConsider); {
		node := yetToConsider[index]
		newYetToConsider := []string{}
		newAlreadyConsidered := []string{}

		for _, n := range yetToConsider {
			if _, ok := lanMap[node][n]; ok {
				newYetToConsider = append(newYetToConsider, n)
			}
		}

		for _, n := range alreadyConsidered {
			if _, ok := lanMap[node][n]; ok {
				newAlreadyConsidered = append(newAlreadyConsidered, n)
			}
		}

		cliques = BronKerbosch(append(currentClique, node), newYetToConsider, newAlreadyConsidered, lanMap, cliques)

		yetToConsider = append(yetToConsider[:index], yetToConsider[index+1:]...)
		alreadyConsidered = append(alreadyConsidered, node)
	}
	return cliques
}

func findMaxCliques(lanMap map[string]map[string]bool) string {
	maxClique := []string{}
	allComputers := []string{}
	for key := range lanMap {
		allComputers = append(allComputers, key)
	}
	cliques := BronKerbosch([]string{}, allComputers, []string{}, lanMap, [][]string{})
	for _, c := range cliques {
		if len(c) > len(maxClique) {
			maxClique = c
		}
	}
	sort.Strings(maxClique)
	return strings.Join(maxClique, ",")
}
