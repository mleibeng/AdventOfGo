package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error in opening input.csv")
	}
	defer file.Close()

	readout := csv.NewReader(file)

	// Idea to start:
	// * split result of csv readout into rules and instructions
	// * index each number according to rules
	// * check if the instruction line is in correct order for printing according to index
	// * count the amount of correct instruction lines
	// * enter at instruction of %2 of amount numbers of instruction line
	// * read out value and add to totalsum
	// * print totalsum
}