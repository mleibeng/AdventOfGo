package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var pl = fmt.Println

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file %v", err)
		return nil, err
	}
	return file, nil
}

func main() {
	file, err := openFile("input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var content string

	for scanner.Scan() {
		content += scanner.Text()
	}

	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := r.FindAllString(content, -1)

	var total int

	for _, match := range matches {

		numbers := strings.TrimPrefix(strings.TrimSuffix(match, ")"), "mul(")
		parts := strings.Split(numbers, ",")

		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		product := num1 * num2
		total += product
	}
	pl(total)
}
