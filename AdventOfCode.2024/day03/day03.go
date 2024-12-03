package day03

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var multiply string = `mul\((\d{1,3}),(\d{1,3})\)`

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
	log.Printf("regex: %s", multiply)
	regex := regexp.MustCompile(multiply)
	fullInput := strings.Join(input, "")
	log.Printf("Full input: %s", fullInput)

	matches := regex.FindAllStringSubmatch(fullInput, -1)
	log.Printf("Matches: %s", matches)
	if matches == nil {
		log.Fatal("no matches found")
	}

	var total int64 = 0
	for _, match := range matches {
		log.Printf("Processing: %s", match[0])
		left, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		total = total + left*right
	}

	log.Print(total)
}

func Part2(input []string) {
	expr := multiply + `|do(?:n't)?\(\)`
	log.Printf("regex: %s", expr)
	regex := regexp.MustCompile(expr)

	fullInput := strings.Join(input, "")
	log.Printf("Full input: %s", fullInput)

	matches := regex.FindAllStringSubmatch(fullInput, -1)
	log.Printf("Matches: %s", matches)
	if matches == nil {
		log.Fatal("no matches found")
	}

	var total int64 = 0
	do := true
	for _, match := range matches {
		log.Printf("Processing: %s", match[0])
		if match[0] == "do()" {
			do = true
			continue
		}
		if match[0] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}

		left, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.ParseInt(match[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		total = total + left*right
	}

	log.Print(total)
}
