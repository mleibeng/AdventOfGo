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

	for index, data := range charArray {
		for innerindex, celldata := range data {
			if celldata == 'X' {
				// Check for left
				if innerindex-1 >= 0 && data[innerindex-1] == 'M' {
					if innerindex-2 >= 0 && data[innerindex-2] == 'A' {
						if innerindex-3 >= 0 && data[innerindex-3] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for up
				if index-1 >= 0 && charArray[index-1][innerindex] == 'M' {
					if index-2 >= 0 && charArray[index-2][innerindex] == 'A' {
						if index-3 >= 0 && charArray[index-3][innerindex] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for up left
				if index-1 >= 0 && innerindex-1 >= 0 && charArray[index-1][innerindex-1] == 'M' {
					if index-2 >= 0 && innerindex-2 >= 0 && charArray[index-2][innerindex-2] == 'A' {
						if index-3 >= 0 && innerindex-3 >= 0 && charArray[index-3][innerindex-3] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for down left
				if index+1 < len(charArray) && innerindex-1 >= 0 && charArray[index+1][innerindex-1] == 'M' {
					if index+2 < len(charArray) && innerindex-2 >= 0 && charArray[index+2][innerindex-2] == 'A' {
						if index+3 < len(charArray) && innerindex-3 >= 0 && charArray[index+3][innerindex-3] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for right
				if innerindex+1 < len(data) && data[innerindex+1] == 'M' {
					if innerindex+2 < len(data) && data[innerindex+2] == 'A' {
						if innerindex+3 < len(data) && data[innerindex+3] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for down
				if index+1 < len(charArray) && charArray[index+1][innerindex] == 'M' {
					if index+2 < len(charArray) && charArray[index+2][innerindex] == 'A' {
						if index+3 < len(charArray) && charArray[index+3][innerindex] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for down right
				if index+1 < len(charArray) && innerindex+1 < len(charArray[index+1]) && charArray[index+1][innerindex+1] == 'M' {
					if index+2 < len(charArray) && innerindex+2 < len(charArray[index+2]) && charArray[index+2][innerindex+2] == 'A' {
						if index+3 < len(charArray) && innerindex+3 < len(charArray[index+3]) && charArray[index+3][innerindex+3] == 'S' {
							xmasCounter++
						}
					}
				}
				// Check for up right
				if index-1 >= 0 && innerindex+1 < len(charArray[index-1]) && charArray[index-1][innerindex+1] == 'M' {
					if index-2 >= 0 && innerindex+2 < len(charArray[index-2]) && charArray[index-2][innerindex+2] == 'A' {
						if index-3 >= 0 && innerindex+3 < len(charArray[index-3]) && charArray[index-3][innerindex+3] == 'S' {
							xmasCounter++
						}
					}
				}
			}
		}
	}
	fmt.Println("XMAS Counter:", xmasCounter)
}