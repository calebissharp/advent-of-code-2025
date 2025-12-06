package main

import (
	_ "embed"
	"testing"
)

var example = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`

//go:embed input.txt
var actual string

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  4277556,
		},
		{
			name:  "actual",
			input: actual,
			want:  4719804927602,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3263827,
		},
		// {
		// 	name:  "actual",
		// 	input: actual,
		// 	want:  ,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
