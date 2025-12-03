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

func findLargestDigit(s string, isLast bool) (digit int, rest string) {
	largestDigit := 0
	largestDigitIndex := 0

	for i, digit := range strings.Split(s, "") {
		if !isLast && i >= len(s)-1 {
			continue
		}

		digit, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}

		if digit > largestDigit {
			largestDigit = digit
			largestDigitIndex = i
		}
	}

	return largestDigit, s[largestDigitIndex+1:]
}

func part1(input string) int {
	totalJoltage := 0

	for bank := range strings.SplitSeq(strings.Trim(input, "\n "), "\n") {
		firstDigit, rest := findLargestDigit(bank, false)
		secondDigit, _ := findLargestDigit(rest, true)

		joltage := firstDigit*10 + secondDigit

		totalJoltage += joltage
	}

	return totalJoltage
}

func part2(input string) int {
	return 0
}
