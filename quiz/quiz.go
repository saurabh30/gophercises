package main

import (
	"flag"
	"fmt"
	"gophercises/quiz/test"
)

func main() {

	fileName := flag.String("file", "test/problems.csv", "display colorized output")
	flag.Parse()
	tQ, cQ := test.TakeTest(*fileName)
	// Open the file
	fmt.Printf("Total Questions : %d, Correct Questions: %d \n", tQ, cQ)
}
