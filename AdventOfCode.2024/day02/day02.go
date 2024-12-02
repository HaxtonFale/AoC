package day02

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
	var safes int = 0
	for _, line := range input {
		report, err := readReport(line)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Report: %d", report)

		if isSafe(report) {
			log.Print("Report safe")
			safes = safes + 1
		} else {
			log.Print("Report unsafe")
		}
	}

	fmt.Println(safes)
}

func Part2(input []string) {
	var safes int = 0
	for _, line := range input {
		var report, err = readReport(line)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Report: %d", report)

		if isSafe(report) {
			log.Print("Report safe")
			safes = safes + 1
		} else {
			log.Print("Testing report with a removed element...")
			length := len(report)
			var cpy []int64
			safe := false
			for i := 0; i < length && !safe; i++ {
				cpy = slices.Delete(slices.Clone(report), i, i+1)

				log.Printf("Updated report: %d", cpy)
				safe = isSafe(cpy)
			}

			if safe {
				log.Print("Updated report safe")
				safes = safes + 1
			}
		}
	}

	fmt.Println(safes)
}

func readReport(line string) ([]int64, error) {
	var report []int64

	for _, str := range strings.Split(line, " ") {
		level, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		report = append(report, level)
	}

	return report, nil
}

func isSafe(report []int64) bool {
	length := len(report) - 1

	increasing := report[0] < report[1]
	if increasing {
		log.Print("Report is increasing.")
	} else {
		log.Print("Report is decreasing.")
	}
	for i := 0; i < length; i++ {
		diff := report[i+1] - report[i]
		if (diff > 0) != increasing {
			log.Printf("Incorrect direction of gradient between %d and %d", report[i], report[i+1])
			return false
		}
		if math.Abs(diff) < 1 {
			log.Printf("Gradient too flat between %d and %d", report[i], report[i+1])
			return false
		}
		if math.Abs(diff) > 3 {
			log.Printf("Gradient too steep between %d and %d", report[i], report[i+1])
			return false
		}
	}

	return true
}
