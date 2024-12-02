package day01

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	advslc "github.com/jkratz55/slices"

	"aoc/math"
)

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
	lefts, rights, err := readInput(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	slices.Sort(lefts)
	slices.Sort(rights)
	log.Printf("Lefts: %d", lefts)
	log.Printf("Rights: %d", rights)

	pairs := advslc.Zip(lefts, rights)
	log.Printf("Pairs: %d", pairs)
	distance := advslc.Reduce(pairs, addDistance, 0)

	fmt.Println(distance)
}

func addDistance(distance int64, values advslc.Pair[int64, int64]) int64 {
	log.Printf("Line: %d, %d", values.First, values.Second)
	diff := math.Abs64(values.First - values.Second)
	log.Printf("Distance: %d", diff)
	return (distance + diff)
}

func Part2(input []string) {
	lefts, rights, err := readInput(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	counts := make(map[int64]int64)
	for _, right := range rights {
		counts[right] = counts[right] + 1
	}

	var score int64 = 0
	for _, left := range lefts {
		score = score + left*counts[left]
	}

	fmt.Println(score)
}

func readInput(input []string) ([]int64, []int64, error) {
	length := len(input)
	lefts := make([]int64, 0, length)
	rights := make([]int64, 0, length)

	for _, line := range input {
		log.Printf("Reading line: %s", line)
		splits := strings.Split(line, "   ")

		log.Printf("Split: %s", splits)

		left, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		rights = append(rights, right)
	}

	return lefts, rights, nil
}
