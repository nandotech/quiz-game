package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	problems := parseLines(lines)
	// Begin correct problem counter
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ?\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{q: line[0], a: line[1]}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
