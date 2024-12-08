package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	userInput, err := os.ReadFile("input-user.txt")
	if err != nil {
		fmt.Println("Skipping benchmark tests for")
		return
	}
	for i := 0; i < b.N; i++ {
		run(false, string(userInput))
	}
}

func BenchmarkPart2(b *testing.B) {
	userInput, err := os.ReadFile("input-user.txt")
	if err != nil {
		fmt.Println("Skipping benchmark tests for")
		panic("")
	}
	for i := 0; i < b.N; i++ {
		run(true, string(userInput))
	}
}
