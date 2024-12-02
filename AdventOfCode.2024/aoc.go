package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"aoc/day01"
	"aoc/day02"
)

func main() {
	var day int
	var part int
	var input int
	var logFile string

	flag.IntVar(&day, "d", 1, "Specify day.")
	flag.IntVar(&part, "p", 0, "Specify part.")
	flag.IntVar(&input, "i", -1, "Specify input file.")
	flag.StringVar(&logFile, "l", "", "Path to log file")
	flag.Parse()

	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	if input < 0 {
		input = part
	}
	var path = fmt.Sprintf("input/d%02dp%d.txt", day, input)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	switch day {
	case 1:
		day01.Solve(part, lines)
	case 2:
		day02.Solve(part, lines)

	default:
		log.Fatal(fmt.Errorf("invalid day: %d", day))
	}
}
