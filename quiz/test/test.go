package test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func TakeTest(fileName string) (tQ, cQ int) {
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	// Iterate through the records
	for {
		// Read each record from csv
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
		var input string
		fmt.Scanln(&input)
		if input == ans {
			cQ++
		}
	}

	return

}
