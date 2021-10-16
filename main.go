package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing file...")
	err := f.Close()
	check(err)
}

func countLinesInFile(f *os.File) int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	// Go back to beginning of file to read questions
	_, _ = f.Seek(0, io.SeekStart)

	return lineCount
}

func main() {
	// Declare -csv and -limit flags
	csvPtr := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "time limit for quiz in seconds")

	// Parse command line input into flags
	flag.Parse()

	fmt.Println("csv file: ", *csvPtr)
	fmt.Println("limit: ", *limitPtr)

	// Actually open problem csv file
	f, err := os.Open(*csvPtr)
	check(err)
	defer closeFile(f)

	countCorrect := 0
	totalQuestionCount := countLinesInFile(f)

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		check(err)

		// Print the question
		fmt.Printf("%s = ", record[0])

		// Parse user answer
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input")
			return
		}

		// Check answer against problem file
		userAnswer := strings.TrimSuffix(input, "\n")
		if userAnswer == record[1] {
			countCorrect += 1
			fmt.Print("✔\n\n")
		} else {
			fmt.Print("✗\n\n")
		}
	}
	fmt.Printf("%d of %d correct.\n", countCorrect, totalQuestionCount)

}
