package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleInput, _ := os.ReadFile("input-example.txt")
	tests := []struct {
		want  any
		name  string
		input string
	}{
		{
			name:  "example",
			input: string(exampleInput),
			want:  int64(37327623),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(false, tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	exampleInput, _ := os.ReadFile("input-example.txt")
	tests := []struct {
		want  any
		name  string
		input string
	}{
		{
			name:  "example",
			input: string(exampleInput),
			want:  24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(true, tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
