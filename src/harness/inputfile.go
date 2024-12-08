package inputfile

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadUserInput(name string) ([]byte, bool) {
	return readFile("inputs", name)
}

func ReadExampleInput(name string) ([]byte, bool) {
	return readFile("examples", name)
}

func readFile(dir string, name string) ([]byte, bool) {
	root, err := findProjectRoot()
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	b, err := os.ReadFile(root + "/data/" + dir + "/2024/" + name + ".txt")
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	if len(b) == 0 {
		fmt.Println("Empty input file")
		return nil, false
	}
	return b, true
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
