package main

import (
	"flag"
	"fmt"
	"gophercises/quiz/test"
)

func main() {
	fileName := flag.String("file", "test/problems.csv", "file name")
	timeDuration := flag.Int("time", 30, "Time duration in seconds.")
	flag.Parse()
	tQ, cQ := test.TakeTest(*fileName, *timeDuration)
	fmt.Printf("Total Questions : %d, Correct Questions: %d \n", tQ, cQ)
}
