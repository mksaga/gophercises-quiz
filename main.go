package main

import (
	"flag"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	csvPtr := flag.String("csv", "problem.csv", "a csv file in format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "time limit for quiz in seconds")

	flag.Parse()

	fmt.Println("csv file: ", *csvPtr)
	fmt.Println("limit: ", *limitPtr)


}
