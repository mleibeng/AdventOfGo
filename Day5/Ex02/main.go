// Idea to start:
// * split result of csv readout into rules and instructions
// * convert to int arrays
// * check if rule is contained in instructions and check if rule is valid
// 	* else break out
// * count the amount of correct instruction lines
// * enter at instruction of / 2 of amount numbers of instruction line
// * read out value and add to totalsum
// * print totalsum

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isCorrectOrder(numbers []int, num1, num2 int) bool {
	index1 := -1;
	index2 := -1;
	for i, num := range numbers {
		if num == num1 {
			index1 = i
		}
		if num == num2 {
			index2 = i
		}
	}
	return index1 != -1 && index2 != -1 && index1 < index2
}

func containsBothNumbers(numbers []int, num1, num2 int) bool {
	index1 := false
	index2 := false
	for _, num := range numbers {
		if num == num1 {
			index1 = true
		}
		if num == num2 {
			index2 = true
		}
	}
	return index1 && index2
}

func swap(numbers []int, num1, num2 int) {
	index1 := -1
	index2 := -1

	for i, num := range numbers {
		if num == num1 {
			index1 = i
		}
		if num == num2 {
			index2 = i
		}
	}
	numbers[index1], numbers[index2] = numbers[index2], numbers[index1]
}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error in opening input.csv")
	}
	defer file.Close()

	readout := bufio.NewScanner(file)

	var rules []string
	var instructions []string

	for readout.Scan() {
		content := readout.Text()
		if (strings.ContainsRune(content, '|')) {
			rules = append(rules, content)
		} else if strings.ContainsRune(content, ',') {
			instructions = append(instructions, content)
		}
	}

	var pairRules [][2]int

	for _, block := range rules {
		parts := strings.Split(block, "|")
		if len(parts) == 2 {
			num1,_ := strconv.Atoi(parts[0])
			num2,_ := strconv.Atoi(parts[1])
			pairRules = append(pairRules, [2]int{num1,num2})
		}
	}

	var instructionValidation [][]int

	for _, line := range instructions {
		parts := strings.Split(line, ",")
		var numbers[]int
		for _, part := range parts {
			num,_ := strconv.Atoi(part)
			numbers = append(numbers, num)
		}
		instructionValidation = append(instructionValidation, numbers)
	}

	var inCorrectInstructionLines [][]int
	// var totalSum int

	for _, numbers := range instructionValidation {
		valid := true
		for _, pair := range pairRules {
			if containsBothNumbers(numbers, pair[0], pair[1]) && !isCorrectOrder(numbers, pair[0], pair[1]) {
				valid = false
				break
			}
		}
		if !valid {
			inCorrectInstructionLines = append(inCorrectInstructionLines, numbers)
		} else {
			continue
		}
	}

	var newCorrectInstructions [][]int
	var totalSum int

	restart := true
	for restart {
		restart = false
		for _, numbers := range inCorrectInstructionLines {
			for i := 0; i < len(numbers); i++ {
				for _, pair := range pairRules {
						if (containsBothNumbers(numbers, pair[0], pair[1]) && !isCorrectOrder(numbers, pair[0], pair[1])) {
							swap(numbers, pair[0], pair[1])
					}
				}
			}
			newCorrectInstructions = append(newCorrectInstructions, numbers)
		}
	}

	// fmt.Println(newCorrectInstructions)

	for _, numbers := range newCorrectInstructions {
		pos := len(numbers) / 2
		numberToAdd := numbers[pos]
		totalSum += numberToAdd
	}

	fmt.Println(totalSum)

}
