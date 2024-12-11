package benchmark

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// RunSolution runs the solution for a specific day
func RunSolution(day Day) ([]string, error) {
	// Check if the binary exists
	binPath := GetPathForBin(day)
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		fmt.Println("Couldn't find executable at", GetPathForBin(day))
		return []string{}, nil
	}

	args := []string{"test", "-bench=.", binPath}

	// Create command
	cmd := exec.Command("go", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	var output []string
	var wg sync.WaitGroup
	wg.Add(2)

	// Handle stderr in a goroutine
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Fprintln(os.Stderr, scanner.Text())
		}
	}()

	// Handle stdout in a goroutine
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Println(line)
			output = append(output, line)
		}
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("command failed: %w", err)
	}

	return output, nil
}

func findProjectRoot() (string, error) {
	// Start from the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree until go.mod is found
	for {
		// Check if go.mod exists in the current directory
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		// Move to parent directory
		parentDir := filepath.Dir(currentDir)

		// If we've reached the filesystem root without finding go.mod, return an error
		if parentDir == currentDir {
			return "", fmt.Errorf("go.mod file not found")
		}

		currentDir = parentDir
	}
}

func GetPathForBin(day Day) string {
	root, err := findProjectRoot()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/src/2024/days/%02d/", root, day.value)
}

// ParseExecTime parses the execution time from the output
func ParseExecTime(output []string, day Day) Timings {
	timings := Timings{
		Day:        day,
		Part1:      nil,
		Part2:      nil,
		TotalNanos: 0,
	}

	for _, line := range output {
		if !strings.Contains(line, "Benchmark") {
			continue
		}
		fmt.Println(line)

		timingStr, nanos, ok := ParseTime(line)
		if !ok {
			fmt.Fprintf(os.Stderr, "Could not parse timings from line: %s\n", line)
			continue
		}

		if strings.Contains(line, "Part1") {
			timingStrCopy := timingStr // Create a copy to avoid referencing loop variable
			timings.Part1 = &timingStrCopy
		} else if strings.Contains(line, "Part2") {
			timingStrCopy := timingStr // Create a copy to avoid referencing loop variable
			timings.Part2 = &timingStrCopy
		}

		timings.TotalNanos += nanos
	}

	return timings
}

func ParseTime(line string) (string, float64, bool) {
	// Regex to match the benchmark format: numbers followed by "ns/op"
	re := regexp.MustCompile(`\s+([\d.]+)\s+ns/op`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		return "", 0, false
	}

	// Parse the nanoseconds value
	nanos, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return "", 0, false
	}

	// Convert to the most appropriate unit and format the string
	var timeStr string
	switch {
	case nanos >= 1_000_000_000: // >= 1s
		timeStr = fmt.Sprintf("%.2fs", nanos/1_000_000_000)
	case nanos >= 1_000_000: // >= 1ms
		timeStr = fmt.Sprintf("%.2fms", nanos/1_000_000)
	case nanos >= 1_000: // >= 1µs
		timeStr = fmt.Sprintf("%.2fµs", nanos/1_000)
	default:
		timeStr = fmt.Sprintf("%.2fns", nanos)
	}

	return timeStr, nanos, true
}
