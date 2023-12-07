package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalCardPointsSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myNumbers, winningNumbers := getNumbers(scanner.Text())
		totalCardPointsSum += countCardPoints(myNumbers, winningNumbers)
	}

	fmt.Println(totalCardPointsSum)
}

func getNumbers(card string) (map[int]bool, []int) {
	myNumbers := make(map[int]bool)
	winningNumbers := make([]int, 0)

	colonIndex := strings.Index(card, ":")
	numbers := strings.Split(card[colonIndex+1:], "|")
	myNumbersStr := strings.TrimSpace(numbers[0])
	winningNumbersStr := strings.TrimSpace(numbers[1])

	for _, num := range strings.Split(myNumbersStr, " ") {
		val, _ := strconv.Atoi(num)
		myNumbers[val] = true
	}

	for _, num := range strings.Split(winningNumbersStr, " ") {
		val, _ := strconv.Atoi(num)
		if val == 0 {
			continue
		}
		winningNumbers = append(winningNumbers, val)
	}
	return myNumbers, winningNumbers
}

func countCardPoints(myNumbers map[int]bool, winningNumbers []int) int {
	cardPoints := 0

	for _, num := range winningNumbers {
		_, ok := myNumbers[num]
		if ok {
			if cardPoints == 0 {
				cardPoints += 1
			} else {
				cardPoints *= 2
			}
		}
	}

	return cardPoints
}
