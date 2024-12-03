package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
)

func main() {
	day := flag.Int("d", 1, "Specify day.")
	part := flag.Int("p", 1, "Specify part.")
	input := flag.Int("i", -1, "Specify input file.")
	logPath := flag.String("l", "", "Path to log file")
	flag.Parse()

	date := time.Now().Format("2006-01-02")

	if *logPath == "" {
		counter := 0
		for {
			counter = counter + 1
			*logPath = fmt.Sprintf("log/d%02dp%d-%s-%03d.log", *day, *part, date, counter)
			_, err := os.Stat(*logPath)
			if errors.Is(err, os.ErrNotExist) {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err := os.Stat("log")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		} else if errors.Is(err, os.ErrNotExist) {
			os.Mkdir("log", 0777)
		}
	}

	logFile, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	writer := io.MultiWriter(logFile, os.Stderr)
	log.SetOutput(writer)

	if *input < 0 {
		*input = *part
	}

	log.Printf("Solving day %d part %d input %d", *day, *part, *input)
	var path = fmt.Sprintf("input/d%02dp%d.txt", *day, *input)
	inputFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	switch *day {
	case 1:
		day01.Solve(*part, lines)
	case 2:
		day02.Solve(*part, lines)
	case 3:
		day03.Solve(*part, lines)

	default:
		log.Fatalf("invalid day: %d", day)
	}
}
