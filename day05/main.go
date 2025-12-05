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
	numFresh := 0

	sections := strings.Split(strings.Trim(input, "\n "), "\n\n")

	rangesStr, ingredientsStr := sections[0], sections[1]

	rangesStrs := strings.Split(rangesStr, "\n")

	ranges := make([][2]int, len(rangesStrs))

	for i, rangeStr := range rangesStrs {
		s := strings.Split(rangeStr, "-")
		start, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		ranges[i] = [2]int{start, end}
	}

	for ingredientStr := range strings.SplitSeq(ingredientsStr, "\n") {
		ingredient, err := strconv.Atoi(ingredientStr)
		if err != nil {
			panic(err)
		}

		for _, rangePair := range ranges {
			if rangePair[0] <= ingredient && rangePair[1] >= ingredient {
				numFresh++
				break
			}
		}
	}

	return numFresh
}

func part2(input string) int {
	return 0
}
