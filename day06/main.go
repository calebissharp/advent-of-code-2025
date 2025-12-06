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

func applyOp(nums []int, op string) (total int) {
	if len(nums) < 2 {
		panic("nums should have at least two numbers")
	}

	total = nums[0]
	switch op {
	case "*":
		for _, n := range nums[1:] {
			total *= n
		}
	case "+":
		for _, n := range nums[1:] {
			total += n
		}
	}

	return
}

func part1(input string) int {
	rows := strings.Split(strings.Trim(input, "\n "), "\n")

	lastRowIdx := len(rows) - 1
	numRows := rows[:lastRowIdx]
	opRow := rows[lastRowIdx]

	ops := strings.Fields(opRow)

	cols := make([][]int, len(ops))

	for _, row := range numRows {
		words := strings.Fields(row)

		for i, word := range words {
			col := cols[i]
			num, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			cols[i] = append(col, num)
		}

	}

	grandTotal := 0

	for i, op := range ops {
		colTotal := 0
		switch op {
		case "*":
			colTotal = 1
			for j := range cols[i] {
				colTotal *= cols[i][j]
			}
		case "+":
			for j := range cols[i] {
				colTotal += cols[i][j]
			}
		}

		grandTotal += colTotal
	}

	return grandTotal
}

// gross
func part2(input string) (grandTotal int) {
	rows := strings.Split(strings.Trim(input, "\n "), "\n")

	width := len(rows[0])
	for _, row := range rows {
		width = max(width, len(row))
	}

	height := len(rows)

	lastRowIdx := len(rows) - 1
	digitRows := rows[:lastRowIdx]
	opRow := rows[lastRowIdx]
	ops := strings.Fields(opRow)

	emptyColumns := make([]int, 0)

	for x := range width {
		isEmpty := true
		for y := range height {
			if rows[y][x] != ' ' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyColumns = append(emptyColumns, x)
		}
	}

	// Manually add last "column"
	emptyColumns = append(emptyColumns, width)

	for problem, emptyColumn := range emptyColumns {
		prevEmptyColumn := -1
		if problem > 0 {
			prevEmptyColumn = emptyColumns[problem-1]
		}

		op := ops[problem]
		problemWidth := emptyColumn - prevEmptyColumn - 1

		cols := make([]int, problemWidth)

		for i := range problemWidth {
			col := prevEmptyColumn + 1 + i
			colDigit := ""
			for _, row := range digitRows {
				digit := string(row[col])
				if digit != " " {
					colDigit += digit
				}
			}
			parsedDigit, err := strconv.Atoi(colDigit)
			if err != nil {
				panic(err)
			}

			cols[i] = parsedDigit
		}

		grandTotal += applyOp(cols, op)
	}

	return
}
