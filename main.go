package main

import (
	"flag"
	"fmt"
)

func main() {
	csvPtr := flag.String("csv", "problem.csv", "a csv file in format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "time limit for quiz in seconds")

	flag.Parse()
	// commandLineArgs := flag.Args()

	fmt.Println("csv file: ", *csvPtr)
	fmt.Println("limit: ", *limitPtr)
}
