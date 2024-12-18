package main

import (
	"bufio"
	"fmt"
	"os"
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

func isStrictlyIncreasingOrDecreasing(nums []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}
		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

func checkSafety(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	if isStrictlyIncreasingOrDecreasing(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		modifiedNums := append([]int{}, nums[:i]...)
		modifiedNums = append(modifiedNums, nums[i+1:]...)
		if isStrictlyIncreasingOrDecreasing(modifiedNums) {
			return true
		}
	}

	return false
}

func main() {
	file, err := openFile("input.csv")
	if err != nil {
		defer file.Close()
	}

	safeCounter := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content := scanner.Text()

		var nums []int
		for _, intContent := range strings.Split(content, " ") {
			num, err := strconv.Atoi(intContent)
			if err != nil {
				pl("Error in strconv atoi function", err)
				continue
			}
			nums = append(nums, num)
		}

		pl("numbers", nums)

		if checkSafety(nums) {
			pl(safeCounter)
			safeCounter++
		}
	}

	pl("result:", safeCounter)
}
