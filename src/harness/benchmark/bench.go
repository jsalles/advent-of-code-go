package benchmark

import (
	"fmt"
	"os"
	"strconv"
)

const (
	ANSI_BOLD   = "\x1b[1m"
	ANSI_RESET  = "\x1b[0m"
	ANSI_ITALIC = "\x1b[3m"
)

func AllDays() []Day {
	days := []Day{}

	entries, err := os.ReadDir("./src/2024/days/")
	if err != nil {
		return days
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dayNum, err := strconv.Atoi(entry.Name())
			if err == nil {
				days = append(days, Day{dayNum})
			}
		}
	}

	return days
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
