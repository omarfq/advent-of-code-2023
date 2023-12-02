package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	result := 0

	// start reading file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += getCalibrationDigits(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func getCalibrationDigits(line string) int {
	left, right := 0, len(line)-1

	digit := [2]int{-1, -1}

	for left <= right {
		if digit[0] == -1 {
			if num, err := strconv.Atoi(string(line[left])); err == nil {
				digit[0] = num
			} else {
				left += 1
			}
		}
		if digit[1] == -1 {
			if num, err := strconv.Atoi(string(line[right])); err == nil {
				digit[1] = num
			} else {
				right -= 1
			}
		}

		if digit[0] != -1 && digit[1] != -1 {
			break
		}
	}

	result := digit[0]*10 + digit[1]
	return result
}
