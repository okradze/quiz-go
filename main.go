package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "total time limit for quiz in seconds")
	flag.Parse()

	lines := readCSVFile(csvFilename)
	problems := parseCSVLines(lines)

	askUserIfReady(timeLimit)
	
	var correctAnswers int

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	go func() {
		<- timer.C
		fmt.Println("\nTime Expired")
		fmt.Printf("You scored %d out of %d\n", correctAnswers, len(lines))
		os.Exit(0)
	}()

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i + 1, problem.question)

		var userAnswer string
		_, err := fmt.Scanf("%s\n", &userAnswer)
		checkError(err)

		if userAnswer == problem.answer {
			correctAnswers++
		}
	}

	fmt.Printf("You scored %d out of %d\n", correctAnswers, len(lines))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func askUserIfReady(limit *int) {
	fmt.Printf("Time limit is %d seconds. Press enter when you are ready.", *limit)
	_, err := fmt.Scanf("\n")
	checkError(err)
}

func readCSVFile(filename *string) [][]string {
	file, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *filename)
		os.Exit(1)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse the CSV file")
		os.Exit(1)
	}

	return lines
}

type problem struct {
	question, answer string
}

func parseCSVLines(lines [][]string) []problem {
	result := make([]problem, len(lines))

	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return result
}