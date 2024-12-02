package day01

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"aoc/math"
)

func Solve(part int, input []string) {
	switch part {
	case 1:
		Part1(input)
	case 2:
		Part2(input)
	default:
		log.Fatal(fmt.Errorf("invalid part: %d", part))
	}
}

func Part1(input []string) {
	length := len(input)
	lefts, rights, err := readInput(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range input {
		splits := strings.Split(line, "   ")

		left, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse value: %s", splits[0]))
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse value: %s", splits[1]))
		}
		rights = append(rights, right)
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	var distance int64 = 0
	for i := 0; i < length; i++ {
		var diff = lefts[i] - rights[i]

		distance = distance + math.Abs(diff)
	}

	fmt.Println(distance)
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
		splits := strings.Split(line, "   ")

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
