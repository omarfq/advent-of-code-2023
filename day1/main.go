package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type digit struct {
	index int
	value string
}

var digitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += searchDigits(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func searchDigits(line string) int {
	var digitsFound []digit

	for k := range digitsMap {
		if strings.Contains(line, k) {
			digitFirstIndex := strings.Index(line, k)
			digitLastIndex := strings.LastIndex(line, k)
			firstDigitFound := digit{
				index: digitFirstIndex,
				value: k,
			}
			lastDigitFound := digit{
				index: digitLastIndex,
				value: k,
			}
			digitsFound = append(digitsFound, firstDigitFound)
			digitsFound = append(digitsFound, lastDigitFound)
		}
	}

	for _, v := range digitsMap {
		strDigit := strconv.Itoa(v)
		if strings.Contains(line, strDigit) {
			digitFirstIndex := strings.Index(line, strDigit)
			digitLastIndex := strings.LastIndex(line, strDigit)
			firstDigitFound := digit{
				index: digitFirstIndex,
				value: strDigit,
			}
			lastDigitFound := digit{
				index: digitLastIndex,
				value: strDigit,
			}
			digitsFound = append(digitsFound, firstDigitFound)
			digitsFound = append(digitsFound, lastDigitFound)
		}
	}

	fmt.Println(digitsFound)
	return getDigitsSum(digitsFound)
}

func getDigitsSum(digitsFound []digit) int {
	digitSum := 0

	sort.Slice(digitsFound, func(i, j int) bool {
		return digitsFound[i].index < digitsFound[j].index
	})

	if firstDigit, err := strconv.Atoi(digitsFound[0].value); err == nil {
		digitSum += firstDigit * 10
	} else {
		digitSum += digitsMap[digitsFound[0].value] * 10
	}

	if lastDigit, err := strconv.Atoi(digitsFound[len(digitsFound)-1].value); err == nil {
		digitSum += lastDigit
	} else {
		digitSum += digitsMap[digitsFound[len(digitsFound)-1].value]
	}

	return digitSum
}

// Solution for Part I
//func getCalibrationDigits(line string) int {
//	left, right := 0, len(line)-1
//
//	digit := [2]int{-1, -1}
//
//	for left <= right {
//		if digit[0] == -1 {
//			if num, err := strconv.Atoi(string(line[left])); err == nil {
//				digit[0] = num
//			} else {
//				left += 1
//			}
//		}
//		if digit[1] == -1 {
//			if num, err := strconv.Atoi(string(line[right])); err == nil {
//				digit[1] = num
//			} else {
//				right -= 1
//			}
//		}
//
//		if digit[0] != -1 && digit[1] != -1 {
//			break
//		}
//	}
//
//	result := digit[0]*10 + digit[1]
//	return result
//}
