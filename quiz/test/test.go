package test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func TakeTest(fileName string, timeDuration int, shuffle bool) (tQ, cQ int) {
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	t := time.NewTimer(time.Duration(timeDuration) * time.Second)

	answerChan := make(chan string)

	var input string

	// fmt.Println("Press enter to start.")
	// fmt.Scanln(&input)

	var problems []Problem

	for {
		// Read each record from csv

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		p := Problem{record[0], record[1]}
		problems = append(problems, p)

	}

	tQ = len(problems)
	if shuffle {
		shuffleProblems(problems)
	}
	for _, problem := range problems {

		go func() {
			fmt.Scanln(&input)
			answerChan <- strings.TrimSpace(input)
		}()

		ans := problem.Answer
		fmt.Printf("Question: %s \n", problem.Question)

		select {
		case <-t.C:
			return

		case answer := <-answerChan:

			if strings.EqualFold(answer, ans) {
				cQ++
			}
		}

	}

	return

}

func shuffleProblems(problems []Problem) {

	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

}
