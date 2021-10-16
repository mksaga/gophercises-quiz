package main

import (
	"flag"
	"fmt"
)

func main() {
	csvPtr := flag.String("csv", "problem.csv", "a csv file in format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "time limit for quiz in seconds")

	flag.Parse()
	commandLineArgs := flag.Args()

	fmt.Println("Hello World!")
	fmt.Println("Args: ", commandLineArgs)
	fmt.Println("csv: ", *csvPtr)
	fmt.Println("limit: ", *limitPtr)
}
