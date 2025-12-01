package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

func part1(input string) int {
	// The number of digits on the dial, starting from 0
	numDigits := 100
	currentDigit := 50

	pw := 0

	for line := range strings.SplitSeq(strings.Trim(input, "\n"), "\n") {
		dir, stepsStr := line[0], line[1:]

		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			panic(err)
		}

		if dir == 'R' {
			currentDigit += steps
		} else {
			currentDigit -= steps
		}

		currentDigit = (currentDigit + numDigits) % numDigits
		if currentDigit == 0 {
			pw += 1
		}
	}

	return pw
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(input string) int {
	// The number of digits on the dial, starting from 0
	numDigits := 100
	currentDigit := 50

	pw := 0

	for line := range strings.SplitSeq(strings.Trim(input, "\n"), "\n") {
		dir, stepsStr := line[0], line[1:]

		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			panic(err)
		}

		newDigit := currentDigit

		if dir == 'R' {
			newDigit += steps
		} else {
			newDigit -= steps
		}

		if newDigit <= 0 || (newDigit%numDigits) == 0 || newDigit > numDigits {
			zerosPassed := max(abs(newDigit/numDigits), 1)
			if currentDigit == 0 && zerosPassed == 1 {
				zerosPassed = 0
			}
			fmt.Println(zerosPassed)
			pw += zerosPassed
		}

		newDigit = (newDigit + numDigits*1000) % numDigits

		fmt.Println(currentDigit, line, newDigit, ":", pw)

		currentDigit = newDigit
	}

	return pw
}
