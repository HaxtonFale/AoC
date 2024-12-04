package day04

import (
	"fmt"
	"log"
	"slices"
	"strings"

	advslc "github.com/jkratz55/slices"
)

var xmas []rune = []rune{'X', 'M', 'A', 'S'}
var samx []rune = []rune{'S', 'A', 'M', 'X'}

func Solve(part int, input []string) {
	switch part {
	case 1:
		Part1(input)
	case 2:
		Part2(input)
	default:
		log.Fatalf("invalid part: %d", part)
	}
}

func Part1(input []string) {
	matrix := advslc.Map(input, func(line string) []rune { return []rune(line) })
	total := 0
	c := make(chan int)

	go horizontal(matrix, c)
	go vertical(matrix, c)
	go diagL(matrix, c)
	go diagR(matrix, c)

	for i := 0; i < 4; i++ {
		count := <-c
		total = total + count
	}

	fmt.Print(total)
}

func Part2(input []string) {
}

func horizontal(input [][]rune, c chan int) {
	log.Print("input:\n", strings.Join(advslc.Map(input, func(line []rune) string { return string(line) }), "\n"))
	count := 0
	for _, row := range input {
		for i := 0; i < len(row)-3; i++ {
			slice := row[i : i+4]
			if slices.Equal(slice, xmas) {
				count++
				continue
			}
			if slices.Equal(slice, samx) {
				count++
			}
		}
	}

	c <- count
}

func vertical(input [][]rune, c chan int) {
	input = transpose(input)
	horizontal(input, c)
}

func diagL(input [][]rune, c chan int) {
	input = tilt(input)
	horizontal(input, c)
}

func diagR(input [][]rune, c chan int) {
	input = tilt(advslc.Map(input, func(row []rune) []rune {
		copy := advslc.Clone(row)
		slices.Reverse(copy)
		return copy
	}))
	horizontal(input, c)
}

func transpose(slice [][]rune) [][]rune {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]rune, xl)

	for i := range result {
		result[i] = make([]rune, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func tilt(slice [][]rune) [][]rune {
	height := len(slice)
	width := len(slice[0])

	limit := min(height, width) - 1

	var output [][]rune
	for i := 0; i < height+width-1; i++ {
		var outputRow []rune
		for j := limit; j >= 0; j-- {
			if (i-j > limit) || (i-j < 0) {
				continue
			}
			outputRow = append(outputRow, slice[j][i-j])
		}
		if len(outputRow) <= 0 {
			continue
		}
		output = append(output, outputRow)
	}
	return output
}
