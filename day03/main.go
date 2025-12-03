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

func findLargestDigit(s string, searchUntilIndex int) (digit int, rest string) {
	largestDigit := 0
	largestDigitIndex := 0

	for i, digit := range strings.Split(s, "") {
		if i >= searchUntilIndex {
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

	if largestDigitIndex > len(s)-1 {
		return largestDigit, ""
	}
	return largestDigit, s[largestDigitIndex+1:]
}

func part1(input string) int {
	totalJoltage := 0

	for bank := range strings.SplitSeq(strings.Trim(input, "\n "), "\n") {
		firstDigit, rest := findLargestDigit(bank, len(bank)-1)
		secondDigit, _ := findLargestDigit(rest, len(bank))

		joltage := firstDigit*10 + secondDigit

		totalJoltage += joltage
	}

	return totalJoltage
}

func part2(input string) int {
	totalJoltage := 0

	bankSize := 12

	for bank := range strings.SplitSeq(strings.Trim(input, "\n "), "\n") {
		joltage := 0

		for position := range bankSize + 1 {
			// how many digits to leave for future positions
			searchUntilIndex := len(bank) - (bankSize - position)
			digit, rest := findLargestDigit(bank, searchUntilIndex)

			positionMultiplier := int(math.Pow(float64(10), float64(bankSize-position-1)))
			joltage += digit * positionMultiplier

			bank = rest
		}

		totalJoltage += joltage
	}

	return totalJoltage
}
