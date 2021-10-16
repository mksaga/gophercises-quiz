package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
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

	limitDuration := time.Duration(*limitPtr) * time.Second

	// my initial approach
	//timer := time.AfterFunc(limitDuration, func() {
	//	fmt.Printf("\nTime's up! %d of %d correct.\n", countCorrect, totalQuestionCount)
	//	os.Exit(0)
	//})
	//defer timer.Stop()

	timer := time.NewTimer(limitDuration)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		check(err)

		// Print the question
		fmt.Printf("%s = ", record[0])

		// Receive the answer on a channel - now input is non-blocking
		answerCh := make(chan string)
		go func () {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		// We got a message from timer channel - stop
		case <-timer.C:
			fmt.Printf("\nTime's up! %d of %d correct.\n", countCorrect, totalQuestionCount)
			return

		// We got an answer from user input channel - process answer
		case answer := <-answerCh:
			if answer == record[1] {
				countCorrect += 1
				fmt.Print("✔\n\n")
			} else {
				fmt.Print("✗\n\n")
			}
		}

	}

	fmt.Printf("\n%d of %d correct.\n", countCorrect, totalQuestionCount)
}
