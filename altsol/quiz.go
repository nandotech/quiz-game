package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	qCsv    string
	timeout int
)

func init() {
	flag.StringVar(&qCsv, "quiz", "problems.csv", "A .csv file with questions and answers.")
	flag.IntVar(&timeout, "timeout", 30, "The time limit for answering questions.")
	flag.Parse()
}

func main() {
	if qCsv == "" {
		log.Fatalln("No CSV file provided.")
	}
	questions := readCsv(qCsv)
	qz := quiz{questions: questions, score: 0}

	resultCh := make(chan quiz)
	go func() {
		qz.ask()
		resultCh <- qz
	}()
	timeoutCh := time.After(time.Duration(timeout) * time.Second)

	select {
	case <-resultCh:
	case <-timeoutCh:
		fmt.Println("Timeout!")
	}
	fmt.Println("Score:", qz.score)
}

type quiz struct {
	questions []question
	score     int
}

type question struct {
	problem string
	result  string
}

func (q *quiz) ask() {
	q.score = 0
	for _, question := range q.questions {
		fmt.Println(question.problem)
		var answer string
		fmt.Scanln(&answer)
		if answer == question.result {
			q.score++
		}
	}
}

func readCsv(p string) []question {
	qfile, err := os.Open(p)
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}
	defer qfile.Close()

	csvrd := csv.NewReader(qfile)
	records, err := csvrd.ReadAll()
	if err != nil {
		log.Fatalf("could not parse .csv: %v\n", err)
	}
	var questions []question
	for _, record := range records {
		q := question{problem: record[0], result: record[1]}
		questions = append(questions, q)
	}
	return questions
}
