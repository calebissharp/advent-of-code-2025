package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

func parseRanges(rangesStr string) [][2]int {
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

	return ranges
}

func combineOverlappingRanges(ranges [][2]int) [][2]int {
	// Sort by left-hand values
	sort.Slice(ranges, func(a, b int) bool {
		return ranges[a][0] < ranges[b][0]
	})

	exclusiveRanges := make([][2]int, 1)
	exclusiveRanges[0] = ranges[0]

	for _, rangePair := range ranges[1:] {
		start, end := rangePair[0], rangePair[1]

		prevRange := exclusiveRanges[len(exclusiveRanges)-1]
		// Check if ranges are overlapping, and if so, combine them
		if prevRange[0] <= start && start <= prevRange[1] {
			prevRange[1] = max(end, prevRange[1])
			exclusiveRanges[len(exclusiveRanges)-1] = prevRange
		} else {
			exclusiveRanges = append(exclusiveRanges, [2]int{start, end})
		}
	}

	return exclusiveRanges
}

func part1(input string) int {
	numFresh := 0

	sections := strings.Split(strings.Trim(input, "\n "), "\n\n")

	rangesStr, ingredientsStr := sections[0], sections[1]

	ranges := parseRanges(rangesStr)

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
	sections := strings.Split(strings.Trim(input, "\n "), "\n\n")

	rangesStr := sections[0]
	ranges := combineOverlappingRanges(parseRanges(rangesStr))

	freshCount := 0
	for _, rangePair := range ranges {
		freshCount += rangePair[1] - rangePair[0] + 1
	}

	return freshCount
}
