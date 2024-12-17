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


}