package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)

	}
}

func part1(input string) (numSplits int) {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")

	width := len(rows[0])

	beams := make([]int, width)
	for i, char := range rows[0] {
		if char == 'S' {
			beams[i] = 1
		}
	}

	for _, row := range rows {
		newBeams := beams

		for i, char := range row {
			if char == '^' && beams[i] == 1 {
				numSplits++
				newBeams[i-1] = 1
				newBeams[i] = 0
				newBeams[i+1] = 1
			}
		}

		chars := []rune(row)
		for i, beam := range newBeams {
			if beam == 1 {
				chars[i] = '|'
			}
		}

		beams = newBeams
	}

	return
}

func part2(input string) (numTimelines int) {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")

	width := len(rows[0])

	beams := make([]int, width)
	for i, char := range rows[0] {
		if char == 'S' {
			beams[i] = 1
		}
	}

	for _, row := range rows {
		newBeams := beams

		for i, char := range row {
			if char == '^' && beams[i] > 0 {
				newBeams[i-1] = beams[i] + newBeams[i-1]
				newBeams[i+1] = beams[i] + newBeams[i+1]
				newBeams[i] = 0
			}
		}

		beams = newBeams
	}

	// Count the number of timelines in the last row
	numTimelines = 0
	for _, beam := range beams {
		if beam > 0 {
			numTimelines += beam
		}
	}

	return
}
