package benchmark

import (
	"fmt"
	"os"
)

const (
	ANSI_BOLD   = "\x1b[1m"
	ANSI_RESET  = "\x1b[0m"
	ANSI_ITALIC = "\x1b[3m"
)

func AllDays() []Day {
	// TODO: make this dynamic
	return []Day{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
		{10},
		{11},
	}
}

func RunBenchmarks() {
	timings := make([]Timings, 0)

	// Assuming AllDays() returns a slice of Day
	for i, day := range AllDays() {
		if i > 0 {
			fmt.Println()
		}

		fmt.Printf("%sDay %d%s\n", ANSI_BOLD, day.IntoInner(), ANSI_RESET)
		fmt.Println("------")

		output, err := RunSolution(day)
		if err != nil {
			fmt.Printf("Error running solution: %v\n", err)
			continue
		}

		if len(output) == 0 {
			fmt.Println("Not solved.")
		} else {
			val := ParseExecTime(output, day)
			timings = append(timings, val)
		}
	}

	var totalNanos float64
	for _, timing := range timings {
		totalNanos += timing.TotalNanos
	}
	totalMillis := totalNanos / 1_000_000

	fmt.Printf("\n%sTotal:%s %s%.2fms%s\n",
		ANSI_BOLD, ANSI_RESET,
		ANSI_ITALIC, totalMillis, ANSI_RESET)

	err := Update(timings, totalMillis)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to update readme with benchmarks.")
	} else {
		fmt.Println("Successfully updated README with benchmarks.")
	}
}
