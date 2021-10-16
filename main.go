package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Declare -csv and -limit flags
	csvPtr := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "time limit for quiz in seconds")

	// Parse command line input into flags
	flag.Parse()

	fmt.Println("csv file: ", *csvPtr)
	fmt.Println("limit: ", *limitPtr)

	// Access the problem csv file
	_, err := os.ReadFile(*csvPtr)
	check(err)

	// Actually open problem csv file
	f, err := os.Open(*csvPtr)
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err!= nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}


}
