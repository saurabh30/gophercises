package test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func TakeTest(fileName string, timeDuration int) (tQ, cQ int) {
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

	for {
		// Read each record from csv
		go func() {
			fmt.Scanln(&input)
			answerChan <- strings.TrimSpace(input)
		}()

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tQ++
		ans := record[1]
		fmt.Printf("Question: %s \n", record[0])

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
