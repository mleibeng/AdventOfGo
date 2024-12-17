package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var pl = fmt.Println

func main() {

	file, err := os.Open("input.csv")
	if err != nil {
		pl("Error in opening input file")
	}
	defer file.Close()

	r := csv.NewReader(file)


	completeReadout, err := r.ReadAll()
	if err != nil {
		pl("Error in readout of file")
		return
	}

	charArray := make([][]byte, len(completeReadout))

	for i, content := range completeReadout {
		if len(content) > 0 {
			charArray[i] = []byte(completeReadout[i][0])
		}
	}


	xmasCounter := 0
/*
   Formats that work:
   M . M | M . S | S . S | S . M |
   . A . | . A . | . A . | . A . |
   S . S | M . S | M . M | S . M |
   use A as pivot point
*/
	for index, data := range charArray {
		for innerindex, celldata := range data {
			if celldata == 'A' {
				// Check for up X
				if index - 1 >= 0 && innerindex - 1 >= 0 && index + 1 < len(charArray) && innerindex+1 < len(charArray[index]) {
					if charArray[index-1][innerindex-1] == 'M' && charArray[index-1][innerindex+1] == 'M' {
						if charArray[index+1][innerindex-1] == 'S' && charArray[index+1][innerindex+1] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for left X
				if index - 1 >= 0 && innerindex - 1 >= 0 && index + 1 < len(charArray) && innerindex+1 < len(charArray[index]) {
					if charArray[index-1][innerindex-1] == 'M' && charArray[index-1][innerindex+1] == 'S' {
						if charArray[index+1][innerindex-1] == 'M' && charArray[index+1][innerindex+1] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for down X
				if index - 1 >= 0 && innerindex - 1 >= 0 && index + 1 < len(charArray) && innerindex+1 < len(charArray[index]) {
					if charArray[index-1][innerindex-1] == 'S' && charArray[index-1][innerindex+1] == 'S' {
						if charArray[index+1][innerindex-1] == 'M' && charArray[index+1][innerindex+1] == 'M' {
							xmasCounter++
						}
					}
				}
				// Check for right X
				if index - 1 >= 0 && innerindex - 1 >= 0 && index + 1 < len(charArray) && innerindex+1 < len(charArray[index]) {
					if charArray[index-1][innerindex-1] == 'S' && charArray[index-1][innerindex+1] == 'M' {
						if charArray[index+1][innerindex-1] == 'S' && charArray[index+1][innerindex+1] == 'M' {
							xmasCounter++
						}
					}
				}
			}
		}
	}
	fmt.Println("XMAS Counter:", xmasCounter)
}