package main

import (
	"fmt"
	"gophercises/quiz/test"
)

func main() {

	tQ, cQ := test.TakeTest("input.csv")
	// Open the file
	fmt.Printf("Total Questions : %d, Correct Questions: %d \n", tQ, cQ)
}
