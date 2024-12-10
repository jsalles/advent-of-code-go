package benchmark

import (
	"fmt"
	"os"
	"strings"
)

// MARKER is used to locate the benchmarking table in the README
const MARKER = "<!--- benchmarking table --->"

// Error represents custom error types
type Error struct {
	Type string
	Msg  string
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Msg)
}

// Day represents a day number
type Day struct {
	value int
}

func (d Day) IntoInner() int {
	return d.value
}

// Timings holds timing information for a day's solutions
type Timings struct {
	Part1      *string
	Part2      *string
	Day        Day
	TotalNanos float64
}

// TablePosition holds the start and end positions of the table in the README
type TablePosition struct {
	posStart int
	posEnd   int
}

// LocateTable finds the position of the benchmarking table in the README
func LocateTable(readme string) (TablePosition, error) {
	matches := strings.Split(readme, MARKER)
	if len(matches) > 3 {
		return TablePosition{}, Error{"Parser", "too many occurences of marker in README"}
	}
	if len(matches) < 2 {
		return TablePosition{}, Error{"Parser", "Could not find table markers"}
	}

	posStart := strings.Index(readme, MARKER)
	posEnd := strings.LastIndex(readme, MARKER) + len(MARKER)

	return TablePosition{posStart, posEnd}, nil
}

// ConstructTable creates the benchmarking table string
func ConstructTable(prefix string, timings []Timings, totalMillis float64) string {
	header := fmt.Sprintf("%s Benchmarks", prefix)
	lines := []string{
		MARKER,
		header,
		"",
		"| Day | Part 1 | Part 2 |",
		"| :---: | :---: | :---:  |",
	}

	for _, timing := range timings {
		part1Val := "-"
		if timing.Part1 != nil {
			part1Val = *timing.Part1
		}
		part2Val := "-"
		if timing.Part2 != nil {
			part2Val = *timing.Part2
		}

		binPath := fmt.Sprintf("./src/2024/days/%02d/code.go", timing.Day.value)
		lines = append(lines, fmt.Sprintf("| [Day %d](%s) | `%s` | `%s` |",
			timing.Day.IntoInner(),
			binPath,
			part1Val,
			part2Val))
	}

	lines = append(lines, "")
	lines = append(lines, fmt.Sprintf("**Total: %.2fms**", totalMillis))
	lines = append(lines, MARKER)

	return strings.Join(lines, "\n")
}

// UpdateContent updates the README content with new timing information
func UpdateContent(s *string, timings []Timings, totalMillis float64) error {
	positions, err := LocateTable(*s)
	if err != nil {
		return err
	}

	table := ConstructTable("##", timings, totalMillis)
	*s = (*s)[:positions.posStart] + table + (*s)[positions.posEnd:]
	return nil
}

// Update updates the README.md file with new timing information
func Update(timings []Timings, totalMillis float64) error {
	path := "README.md"
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	readme := string(content)
	if err := UpdateContent(&readme, timings, totalMillis); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(readme), 0644)
}
