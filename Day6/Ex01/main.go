package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

/*
	BattlePlan:
	* read in map
	* find position of guard and heading of guard
	* program moves for guard in map:
		[-1 ][ 0 ] == up
		[+1 ][ 0 ] == down
		[ 0 ][-1 ] == left
		[ 0 ][+1 ] == right
	if guard hits # then -> rotate to next command.
		if heading == up start with up -> right -> down -> left
		etc...
	previous guard position convert to -> X after every move
	if guard next move -> out of bounds break out of loop. and convert last pos to X as well!!
	finally after guard is out -> count X's on map.
*/

var (
	up = []int{-1,0}
	down = []int{1,0}
	left = []int{0,-1}
	reight = []int{0,1}
)

func move(field [][]string, guardPos [2]int, direction []int) ([2]int, error) {

	nextGuardPos := [2]int{guardPos[0] + direction[0], guardPos[1] + direction[1]}

	if (nextGuardPos[0] < 0 || nextGuardPos[0] >= len(field) || nextGuardPos[1] < 0 || nextGuardPos[1] >= len(field[0])) {
		return guardPos, fmt.Errorf("move out of bounds")
	}
	if (field[nextGuardPos[0]][nextGuardPos[1]] == "#") {

	}
	return nextGuardPos, nil
}

func isMoveBlocked(field [][]string, nextGuardPos [2]int, direction []int)

func findGuardPos(field [][]string) ([2]int, error) {

	var guardPos [2]int = [2]int{-1,-1}

	for i, row := range field {
		for j, cell := range row {
			if cell == "^" {
				guardPos = [2]int{i,j}
				return guardPos, nil
			}
		}
	}
	if guardPos == [2]int{-1,-1} {
		return guardPos, errors.New("no guard found")
	}
	return guardPos, nil
}

func main() {
	file, err := os.Open("input2.csv")
	if err != nil {
		fmt.Println("Error in opening file")
	}
	defer file.Close()

	readout := csv.NewReader(file)

	field, err := readout.ReadAll()
	if err != nil {
		fmt.Println("Error in readout")
	}

	guardPos, err := findGuardPos(field)
	if err != nil {
		fmt.Println(err)
	}

	/*for {
	condition 1 = check which direction to go
	condition 2 = check if next position would be blocked
	condition 3 rotate 90degress
	condition 4 do the move
	condition 5 check if move went out of map
		if yes -> update last position with X && end loop
		if no -> update last position with X and update to new position
	}
	*/

	for {

	}

	// count X's

}