package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// parse flag names
	csvFilename := flag.String("csv", "problems.csv",
		"csv file in format of qna")
	// solveTime := flag.Float64("time", 3.0, "time limit to solve problems")
	numberTries := flag.Int("tries", 1, "number of tries to answer question")
	if *numberTries == 0 {
		*numberTries = 1 // set it to default!!! if 0!!
	}
	flag.Parse()

	// open csv then close it
	f, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("Failed to open csv file: %s\n", csvFilename)
		os.Exit(1)
	}
	defer f.Close()

	// read file into data.
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Failed to open csv file: %s\n", csvFilename)
	}

	// the quiz itself
	var answer string
	for _, val := range data {
		fmt.Printf("Answer all these questions and you will win: \n")
		fmt.Println(val[0])
		fmt.Scanf("%s", &answer)
		if answer == val[1] {
			fmt.Println("Woah you guessed right! Very nice. \n")
		} else {
			fmt.Println("WRONG!! To the gulag with you. \n")
		}
	}
}
