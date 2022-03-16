package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

/* Learned how to:
- Use a basic goroutine: select/go (some minor concurrency)
- Parse CLI arguments
- Open/read/parse csv files
- Go label, but ymm
- TODO: what is a go closure.
- TODO: I could refactor
*/

func main() {
	// parse flag names
	csvFilename := flag.String("csv", "problems.csv", "csv file in format of qna")
	solveTime := flag.Int("time", 30.0, "time limit to solve the entire quiz in seconds")
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
		fmt.Println("Failed to read csv file: %s\n", csvFilename)
	}

	// the quiz itself
	correctAnswers := 0

	fmt.Printf("Answer all these questions and you will win (knowledge). \n")
	quizTimer := time.NewTimer(time.Duration(*solveTime) * time.Second)
	// Only start the quiz if the user has pressed any key
	var answer string
problemLoop:
	for _, val := range data {
		fmt.Println(val[0])
		answerChannel := make(chan string)
		go func() { // scans for an answer
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer // we send the input/answer to the channel.
		}()

		select {
		// we remove the default
		// because even if we run out of time and type an answer,
		// it is counted as a correct answer (?)
		case <-quizTimer.C: // if the timer ran out of time.
			break problemLoop
		case answer := <-answerChannel: // if we get a call from the answerChannel before the timer
			if answer == val[1] {
				correctAnswers++
				fmt.Println("Woah you guessed right! Very nice. \n")
			} else {
				fmt.Println("WRONG!! To the gulag with you. \n")
			}
		}
	}
	fmt.Printf("You have answered %d questions correctly and %d wrong! \n", correctAnswers, len(data)-correctAnswers)
}
