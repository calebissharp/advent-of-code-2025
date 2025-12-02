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
	sum := 0

	for rangeStr := range strings.SplitSeq(strings.Trim(input, "\n "), ",") {
		split := strings.Split(rangeStr, "-")

		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			productID := strconv.Itoa(i)
			if productID[0:len(productID)/2] == productID[len(productID)/2:] {
				sum += i
			}
		}

	}
	return sum
}

func part2(input string) int {
	return 100
}
