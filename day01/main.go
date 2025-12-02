package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mod(n int, q int) int {
	return ((n % q) + q) % q
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

		currentDigit = mod(currentDigit, numDigits)
		if currentDigit == 0 {
			pw += 1
		}
	}

	return pw
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

		op := 0

		if dir == 'R' {
			op = steps
		} else {
			op = -steps
		}

		// NOTE: Need to use float math because negative integer division rounds *up* (ie. towards 0)
		numRots := math.Abs(math.Floor(float64(currentDigit+op) / float64(numDigits)))

		newDigit := mod(currentDigit+op, numDigits)

		// FIXME: special cases for left-handed rotations that start or end on 0
		// Thre's probably a better way of handling this but I'm too stupid. It's
		// probably some kind of off-by-one error in the logic
		if op < 0 {
			if currentDigit == 0 {
				numRots = max(0, numRots-1)
			}
			if newDigit == 0 {
				numRots += 1
			}
		}

		pw += int(numRots)

		currentDigit = newDigit
	}

	return pw
}
