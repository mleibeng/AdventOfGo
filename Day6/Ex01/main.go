package main

import (
	"bufio"
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
	up = [2]int{-1,0}
	down = [2]int{1,0}
	left = [2]int{0,-1}
	right = [2]int{0,1}
)

var directions = [][]int{
	{up[0], up[1]},
	{right[0], right[1]},
	{down[0], down[1]},
	{left[0], left[1]},
}

func move(field [][]string, guardPos [2]int, direction [2]int) ([2]int, error) {

	nextGuardPos := [2]int{guardPos[0] + direction[0], guardPos[1] + direction[1]}

	if (nextGuardPos[0] < 0 || nextGuardPos[0] >= len(field) || nextGuardPos[1] < 0 || nextGuardPos[1] >= len(field[0])) {
		return guardPos, fmt.Errorf("move out of bounds")
	}
	return nextGuardPos, nil
}

func isMoveBlocked(field [][]string, guardPos [2]int, direction [2]int) (bool) {

	nextGuardPos := [2]int{guardPos[0] + direction[0], guardPos[1] + direction[1]}

	if (nextGuardPos[0] < 0 || nextGuardPos[0] >= len(field) || nextGuardPos[1] < 0 || nextGuardPos[1] >= len(field[0])) {
		return false
	}

	if (field[nextGuardPos[0]][nextGuardPos[1]] == "#") {
		return true
	}
	return false
}

func changeDirection(currentDirection [2]int) [2]int {
	for i, dir := range directions {
		if dir[0] == currentDirection[0] && dir[1] == currentDirection[1] {
			return [2]int{directions[(i+1)%len(directions)][0], directions[(i+1)%len(directions)][1]}
		}
	}
	return currentDirection
}

func splitCharacters(line string) []string {

	var characters []string
	for _, char := range line {
		characters = append(characters, string(char))
	}

	return characters
}

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
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error in opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var field [][]string

	for scanner.Scan() {
		line := scanner.Text()
		characters := splitCharacters(line)
		field = append(field, characters)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error in scanner")
		return
	}

	guardPos, err := findGuardPos(field)
	if err != nil {
		fmt.Println(err)
	}

	/*for {
	condition 1 = check if next position would be blocked
		condition 2 rotate 90degress
		condition 3 check if also blocked
		etc.. until unblocked

	condition 4 do the move
	condition 5 check if move went out of map
		if yes -> update last position with X && end loop
		if no -> update last position with X and update to new position
	}
	*/

	currentDirection := up;
	oldPos := guardPos
	var visitCounter int

	for {
		if (isMoveBlocked(field, oldPos, currentDirection)) {
			currentDirection = changeDirection(currentDirection)
		} else {
			newPos, err := move(field, oldPos, currentDirection)
			if err != nil {
				visitCounter++
				fmt.Println(err)
				break ;
			}
			field[oldPos[0]][oldPos[1]] = "X"
			oldPos = newPos
		}
	}


	for _, visited := range field {
		for _,visitedCell := range visited {
			if (visitedCell == "X") {
				visitCounter++
			}
		}
	}
	fmt.Println(visitCounter)

}