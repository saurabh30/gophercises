package main

import (
	"flag"
	"fmt"
	"gophercises/quiz/test"
)

func main() {
	fileName := flag.String("file", "test/problems.csv", "file name")
	timeDuration := flag.Int("time", 30, "Time duration in seconds.")
	shuffle := flag.Bool("shuffle", false, "Boolean flag to shuffle the quiz order.")
	flag.Parse()
	tQ, cQ := test.TakeTest(*fileName, *timeDuration, *shuffle)
	fmt.Printf("Total Questions : %d, Correct Questions: %d \n", tQ, cQ)
}
