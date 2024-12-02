package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "aoc/day01"
)

func main() {
    var day int
    var part int
    var input int

    flag.IntVar(&day, "d", 1, "Specify day.")
    flag.IntVar(&part, "p", 0, "Specify part.")
    flag.IntVar(&input, "i", -1, "Specify input file.")
    flag.Parse()

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

    switch day {
        case 1: day01.Solve(part, lines)

        default: fmt.Errorf("Invalid day: %d", day)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
