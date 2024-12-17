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
	// * index each number,
	// * check if the index ordering related to the first part of the csv is in correct order for the printing orders in the below instructions
	// * count the amount of correct instruction lines
	// * enter at instruction of %2 of amount numbers of instruction line
	// * read out value and add to totalsum
	// * print totalsum
}