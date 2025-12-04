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

func part1(input string) int {
	count, _ := removeAccessiblePaper(input)
	return count
}

func removeAccessiblePaper(input string) (count int, newGrid string) {
	rollChar := '@'
	emptyChar := '.'

	grid := strings.Split(strings.Trim(input, "\n "), "\n")

	for y, row := range grid {
		newRow := ""

		for x, char := range row {
			if rune(char) == emptyChar {
				newRow += "."
				continue
			}

			adjacentIndices := [][2]int{
				// Above
				{-1, -1},
				{0, -1},
				{1, -1},
				// left + right
				{-1, 0},
				{1, 0},
				// Below
				{-1, 1},
				{0, 1},
				{1, 1},
			}

			adjacentRolls := 0

			for _, coords := range adjacentIndices {
				xOff, yOff := coords[0], coords[1]

				xAdj := x + xOff
				yAdj := y + yOff

				if xAdj < 0 || xAdj >= len(row) || yAdj < 0 || yAdj >= len(grid) {
					continue
				}

				adjacentChar := grid[yAdj][xAdj]

				if rune(adjacentChar) == rollChar {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				count++
				newRow += "."
			} else {
				newRow += string(char)
			}
		}

		newGrid += newRow + "\n"
	}

	return
}

func part2(input string) int {
	totalCount := 0

	for {
		count, newGrid := removeAccessiblePaper(input)

		totalCount += count

		if count == 0 {
			return totalCount
		}

		input = newGrid
	}
}
